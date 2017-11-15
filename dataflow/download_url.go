/*
 * Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
 *
 * This program and the accompanying materials are made available under
 * the terms of the under the Apache License, Version 2.0 (the "License”);
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package dataflow

import "github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/httpclient"

func DataflowShellDownloadUrl(dataflowServer string, authClient httpclient.AuthenticatedClient) (string, string, error) {
	return "https://repo.spring.io/libs-snapshot/org/springframework/cloud/spring-cloud-dataflow-shell/1.2.3.RELEASE/spring-cloud-dataflow-shell-1.2.3.RELEASE.jar",
		"9dec3eab5740cb087d7842bcb6bf924f9e008638dedeca16c5336bbc3c0e4453", nil
}
