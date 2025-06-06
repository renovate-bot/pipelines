// Copyright 2018-2023 The Kubeflow Authors
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

package integration

import (
	"os"
	"testing"

	"github.com/golang/glog"
	api_server "github.com/kubeflow/pipelines/backend/src/common/client/api_server/v2"
	test "github.com/kubeflow/pipelines/backend/test/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HealthzApiTest struct {
	suite.Suite
	namespace     string
	healthzClient *api_server.HealthzClient
}

// Check the namespace have ML job installed and ready
func (s *HealthzApiTest) SetupTest() {
	if !*runIntegrationTests {
		s.T().SkipNow()
		return
	}

	if !*isDevMode {
		err := test.WaitForReady(*initializeTimeout)
		if err != nil {
			glog.Exitf("Failed to initialize test. Error: %v", err)
		}
	}
	s.namespace = *namespace
	clientConfig := test.GetClientConfig(*namespace)
	var err error
	s.healthzClient, err = api_server.NewHealthzClient(clientConfig, false)
	if err != nil {
		glog.Exitf("Failed to get healthz client. Error: %v", err)
	}
	s.cleanUp()
}

func (s *HealthzApiTest) TearDownSuite() {
	if *runIntegrationTests {
		if !*isDevMode {
			s.cleanUp()
		}
	}
}

func (s *HealthzApiTest) cleanUp() {
}

func (s *HealthzApiTest) TestHealthzAPI() {
	t := s.T()

	/* ---------- Verify healthz response ---------- */
	healthzResp, err := s.healthzClient.GetHealthz()
	assert.Nil(t, err)
	assert.NotNil(t, healthzResp)

	pipelineStore := os.Getenv("PIPELINE_STORE")
	if pipelineStore == "" {
		pipelineStore = "database"
	}

	assert.Equal(t, pipelineStore, healthzResp.PipelineStore)
}

func TestHealthzAPI(t *testing.T) {
	suite.Run(t, new(HealthzApiTest))
}
