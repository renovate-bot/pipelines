/*
 * Copyright 2019 The Kubeflow Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as React from 'react';
import PipelineVersionList, { PipelineVersionListProps } from './PipelineVersionList';
import TestUtils from 'src/TestUtils';
import { V2beta1PipelineVersion } from 'src/apisv2beta1/pipeline';
import { Apis, ListRequest } from 'src/lib/Apis';
import { shallow, ReactWrapper, ShallowWrapper } from 'enzyme';
import { range } from 'lodash';

class PipelineVersionListTest extends PipelineVersionList {
  public _loadPipelineVersions(request: ListRequest): Promise<string> {
    return super._loadPipelineVersions(request);
  }
}

describe('PipelineVersionList', () => {
  let tree: ReactWrapper | ShallowWrapper;

  const listPipelineVersionsSpy = jest.spyOn(Apis.pipelineServiceApiV2, 'listPipelineVersions');
  const onErrorSpy = jest.fn();

  function generateProps(): PipelineVersionListProps {
    return {
      history: {} as any,
      location: { search: '' } as any,
      match: '' as any,
      onError: onErrorSpy,
      pipelineId: 'pipeline',
    };
  }

  async function mountWithNPipelineVersions(n: number): Promise<ReactWrapper> {
    listPipelineVersionsSpy.mockImplementation((pipelineId: string) => ({
      pipeline_versions: range(n).map(i => ({
        pipeline_version_id: 'test-pipeline-version-id' + i,
        display_name: 'test pipeline version name' + i,
      })),
    }));
    tree = TestUtils.mountWithRouter(<PipelineVersionList {...generateProps()} />);
    await listPipelineVersionsSpy;
    await TestUtils.flushPromises();
    tree.update(); // Make sure the tree is updated before returning it
    return tree;
  }

  beforeEach(() => {
    jest.clearAllMocks();
  });

  afterEach(async () => {
    // unmount() should be called before resetAllMocks() in case any part of the unmount life cycle
    // depends on mocks/spies
    await tree.unmount();
    jest.resetAllMocks();
  });

  it('renders an empty list with empty state message', () => {
    tree = shallow(<PipelineVersionList {...generateProps()} />);
    expect(tree).toMatchSnapshot();
  });

  it('renders a list of one pipeline version', async () => {
    tree = shallow(<PipelineVersionList {...generateProps()} />);
    tree.setState({
      pipelineVersions: [
        {
          created_at: new Date(2018, 8, 22, 11, 5, 48),
          display_name: 'pipelineversion1',
          name: 'pipelineversion1',
        } as V2beta1PipelineVersion,
      ],
    });
    await listPipelineVersionsSpy;
    expect(tree).toMatchSnapshot();
  });

  it('renders a list of one pipeline version with description', async () => {
    tree = shallow(<PipelineVersionList {...generateProps()} />);
    tree.setState({
      pipelineVersions: [
        {
          created_at: new Date(2018, 8, 22, 11, 5, 48),
          display_name: 'pipelineversion1',
          name: 'pipelineversion1',
          description: 'pipelineversion1 description',
        } as V2beta1PipelineVersion,
      ],
    });
    await listPipelineVersionsSpy;
    expect(tree).toMatchSnapshot();
  });

  it('renders a list of one pipeline version without created date', async () => {
    tree = shallow(<PipelineVersionList {...generateProps()} />);
    tree.setState({
      pipelines: [
        {
          display_name: 'pipelineversion1',
          name: 'pipelineversion1',
        } as V2beta1PipelineVersion,
      ],
    });
    await listPipelineVersionsSpy;
    expect(tree).toMatchSnapshot();
  });

  it('renders a list of one pipeline version with error', async () => {
    tree = shallow(<PipelineVersionList {...generateProps()} />);
    tree.setState({
      pipelineVersions: [
        {
          created_at: new Date(2018, 8, 22, 11, 5, 48),
          error: 'oops! could not load pipeline',
          display_name: 'pipeline1',
          name: 'pipeline1',
          parameters: [],
        } as V2beta1PipelineVersion,
      ],
    });
    await listPipelineVersionsSpy;
    expect(tree).toMatchSnapshot();
  });

  it('calls Apis to list pipeline versions, sorted by creation time in descending order', async () => {
    tree = await mountWithNPipelineVersions(2);
    await (tree.instance() as PipelineVersionListTest)._loadPipelineVersions({
      pageSize: 10,
      pageToken: '',
      sortBy: 'created_at',
    } as ListRequest);
    expect(listPipelineVersionsSpy).toHaveBeenLastCalledWith(
      'pipeline',
      '',
      10,
      'created_at',
      undefined,
    );
    expect(tree).toMatchSnapshot();
  });

  it('calls Apis to list pipeline versions, sorted by pipeline version name in descending order', async () => {
    tree = await mountWithNPipelineVersions(3);
    await (tree.instance() as PipelineVersionListTest)._loadPipelineVersions({
      pageSize: 10,
      pageToken: '',
      sortBy: 'name',
    } as ListRequest);
    expect(listPipelineVersionsSpy).toHaveBeenLastCalledWith('pipeline', '', 10, 'name', undefined);
    expect(tree).toMatchSnapshot();
  });
});
