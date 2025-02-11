/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package inst

import (
	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
	// test "github.com/openark/golib/tests"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func TestGetAnalysisInstanceType(t *testing.T) {
	// {
	// 	analysis := &ReplicationAnalysis{}
	// 	test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "intermediate-master")
	// }
	// {
	// 	analysis := &ReplicationAnalysis{IsMaster: true}
	// 	test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "master")
	// }
	// {
	// 	analysis := &ReplicationAnalysis{IsCoMaster: true}
	// 	test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "co-master")
	// }
	// {
	// 	analysis := &ReplicationAnalysis{IsMaster: true, IsCoMaster: true}
	// 	test.S(t).ExpectEquals(string(analysis.GetAnalysisInstanceType()), "co-master")
	// }
}
