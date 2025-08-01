# Copyright 2024 The Kubeflow Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import os

from kfp.dsl import component

PACKAGES_TO_INSTALL = ['yapf']
if 'KFP_PIPELINE_SPEC_PACKAGE_PATH' in os.environ:
    PACKAGES_TO_INSTALL.append(os.environ['KFP_PIPELINE_SPEC_PACKAGE_PATH'])


@component(
    pip_index_urls=['https://pypi.org/simple'],
    packages_to_install=['yapf'],
    use_venv=True,
)
def component_with_pip_install():
    import yapf

    print(dir(yapf))


if __name__ == '__main__':
    from kfp import compiler

    compiler.Compiler().compile(
        pipeline_func=component_with_pip_install,
        package_path=__file__.replace('.py', '.yaml'),
    )
