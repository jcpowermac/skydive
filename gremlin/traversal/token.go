/*
 * Copyright (C) 2018 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package traversal

import "github.com/skydive-project/skydive/topology/graph/traversal"

const (
	traversalFlowToken        traversal.Token = 1001
	traversalHopsToken        traversal.Token = 1002
	traversalNodesToken       traversal.Token = 1003
	traversalCaptureNodeToken traversal.Token = 1004
	traversalAggregatesToken  traversal.Token = 1005
	traversalRawPacketsToken  traversal.Token = 1006
	traversalBpfToken         traversal.Token = 1007
	traversalMetricsToken     traversal.Token = 1008
	traversalSocketsToken     traversal.Token = 1009
)
