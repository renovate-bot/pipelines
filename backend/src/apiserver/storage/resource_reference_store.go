// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"k8s.io/apimachinery/pkg/util/json"
)

var resourceReferenceColumns = []string{
	"ResourceUUID", "ResourceType", "ReferenceUUID",
	"ReferenceName", "ReferenceType", "Relationship", "Payload",
}

type ResourceReferenceStoreInterface interface {
	// Retrieve the resource reference for a given resource id, type and a reference type.
	GetResourceReference(resourceId string, resourceType model.ResourceType,
		referenceType model.ResourceType) (*model.ResourceReference, error)
}

type ResourceReferenceStore struct {
	db            *DB
	pipelineStore PipelineStoreInterface
}

// Create a resource reference.
// This is always in company with creating a parent resource so a transaction is needed as input.
func (s *ResourceReferenceStore) CreateResourceReferences(tx *sql.Tx, refs []*model.ResourceReference) error {
	if len(refs) > 0 {
		resourceRefSqlBuilder := sq.
			Insert("resource_references").
			Columns("ResourceUUID", "ResourceType", "ReferenceUUID", "ReferenceName", "ReferenceType", "Relationship", "Payload")
		for _, ref := range refs {
			if !s.checkReferenceExist(tx, ref.ReferenceUUID, ref.ReferenceType) {
				return util.NewResourceNotFoundError(string(ref.ReferenceType), ref.ReferenceUUID)
			}
			payload, err := json.Marshal(ref)
			if err != nil {
				return util.NewInternalServerError(err, "Failed to stream resource reference model to a json payload")
			}
			resourceRefSqlBuilder = resourceRefSqlBuilder.Values(
				ref.ResourceUUID, ref.ResourceType, ref.ReferenceUUID, ref.ReferenceName, ref.ReferenceType, ref.Relationship, string(payload))
		}
		refSql, refArgs, err := resourceRefSqlBuilder.ToSql()
		if err != nil {
			return util.NewInternalServerError(err, "Failed to create query to store resource references")
		}
		_, err = tx.Exec(refSql, refArgs...)
		if err != nil {
			return util.NewInternalServerError(err, "Failed to store resource references")
		}
	}
	return nil
}

func (s *ResourceReferenceStore) checkReferenceExist(tx *sql.Tx, referenceId string, referenceType model.ResourceType) bool {
	var selectBuilder sq.SelectBuilder
	switch referenceType {
	case model.JobResourceType:
		selectBuilder = sq.Select("1").From("jobs").Where(sq.Eq{"uuid": referenceId})
	case model.ExperimentResourceType:
		selectBuilder = sq.Select("1").From("experiments").Where(sq.Eq{"uuid": referenceId})
	case model.PipelineVersionResourceType:
		if s.pipelineStore != nil {
			pv, _ := s.pipelineStore.GetPipelineVersion(referenceId)

			return pv != nil
		}
		selectBuilder = sq.Select("1").From("pipeline_versions").Where(sq.Eq{"uuid": referenceId})
	case model.PipelineResourceType:
		if s.pipelineStore != nil {
			p, _ := s.pipelineStore.GetPipeline(referenceId)

			return p != nil
		}
		selectBuilder = sq.Select("1").From("pipelines").Where(sq.Eq{"uuid": referenceId})
	case model.NamespaceResourceType:
		// This function is called to check the data validity when the data are transformed according to the DB schema.
		// Since there is not a separate table to store the namespace data, thus always returning true.
		return true
	default:
		return false
	}
	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return false
	}
	var exists bool
	err = tx.QueryRow(fmt.Sprintf("SELECT exists (%s)", query), args...).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return exists
}

// Delete all resource references for a specific resource.
// This is always in company with creating a parent resource so a transaction is needed as input.
func (s *ResourceReferenceStore) DeleteResourceReferences(tx *sql.Tx, id string, resourceType model.ResourceType) error {
	refSql, refArgs, err := sq.
		Delete("resource_references").
		Where(sq.Or{
			sq.Eq{"ResourceUUID": id, "ResourceType": resourceType},
			sq.Eq{"ReferenceUUID": id, "ReferenceType": resourceType},
		}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete resource references for %s %s due to SQL syntax error", resourceType, id)
	}
	_, err = tx.Exec(refSql, refArgs...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete resource references for %s %s due to SQL execution error", resourceType, id)
	}
	return nil
}

func (s *ResourceReferenceStore) GetResourceReference(resourceId string, resourceType model.ResourceType,
	referenceType model.ResourceType,
) (*model.ResourceReference, error) {
	sql, args, err := sq.Select(resourceReferenceColumns...).
		From("resource_references").
		Where(sq.Eq{
			"ResourceUUID":  resourceId,
			"ResourceType":  resourceType,
			"ReferenceType": referenceType,
		}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err,
			"Failed to create query to get resource reference. "+
				"Resource ID: %s. Resource Type: %s. Reference Type: %s", resourceId, resourceType, referenceType)
	}
	row, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err,
			"Failed to get resource reference. "+
				"Resource ID: %s. Resource Type: %s. Reference Type: %s", resourceId, resourceType, referenceType)
	}
	defer row.Close()
	reference, err := s.scanRows(row)
	if err != nil || len(reference) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get resource reference: %v", err.Error())
	}
	if len(reference) == 0 {
		return nil, util.NewResourcesNotFoundError(
			"Resource ID: %s. Resource Type: %s. Reference Type: %s", resourceId, resourceType, referenceType)
	}
	return &reference[0], nil
}

func (s *ResourceReferenceStore) scanRows(r *sql.Rows) ([]model.ResourceReference, error) {
	var references []model.ResourceReference
	for r.Next() {
		var resourceUUID, resourceType, referenceUUID, referenceName, referenceType, relationship, payload string
		err := r.Scan(
			&resourceUUID, &resourceType, &referenceUUID, &referenceName, &referenceType, &relationship, &payload)
		if err != nil {
			return nil, err
		}
		references = append(references, model.ResourceReference{
			ResourceUUID:  resourceUUID,
			ResourceType:  model.ResourceType(resourceType),
			ReferenceUUID: referenceUUID,
			ReferenceName: referenceName,
			ReferenceType: model.ResourceType(referenceType),
			Relationship:  model.Relationship(relationship),
			Payload:       payload,
		})
	}
	return references, nil
}

func NewResourceReferenceStore(db *DB, pipelineStore PipelineStoreInterface) *ResourceReferenceStore {
	// If pipelineStore is specified and it is not nil, it will be used instead of the DB.
	// This will make pipelines and pipeline versions to get stored in K8s instead of Database.
	return &ResourceReferenceStore{db: db, pipelineStore: pipelineStore}
}
