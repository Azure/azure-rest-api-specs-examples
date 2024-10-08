package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetScriptExecutionHistory.json
func ExampleScriptExecutionHistoryClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScriptExecutionHistoryClient().NewListByClusterPager("rg1", "cluster1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ScriptActionExecutionHistoryList = armhdinsight.ScriptActionExecutionHistoryList{
		// 	Value: []*armhdinsight.RuntimeScriptActionDetail{
		// 		{
		// 			Name: to.Ptr("Test"),
		// 			ApplicationName: to.Ptr("app1"),
		// 			Roles: []*string{
		// 				to.Ptr("headnode"),
		// 				to.Ptr("workernode")},
		// 				URI: to.Ptr("http://testurl.com/install.ssh"),
		// 				EndTime: to.Ptr("2017-03-22T21:34:39.293"),
		// 				ExecutionSummary: []*armhdinsight.ScriptActionExecutionSummary{
		// 				},
		// 				Operation: to.Ptr("PostClusterCreateScriptActionRequest"),
		// 				ScriptExecutionID: to.Ptr[int64](391145124054712),
		// 				StartTime: to.Ptr("2017-03-22T21:34:39.293"),
		// 				Status: to.Ptr("ValidationFailed"),
		// 			},
		// 			{
		// 				Name: to.Ptr("Test"),
		// 				ApplicationName: to.Ptr("app2"),
		// 				Roles: []*string{
		// 					to.Ptr("headnode"),
		// 					to.Ptr("workernode")},
		// 					URI: to.Ptr("http://testurl.com/install-script.ssh"),
		// 					EndTime: to.Ptr("2017-03-22T21:34:39.293"),
		// 					ExecutionSummary: []*armhdinsight.ScriptActionExecutionSummary{
		// 					},
		// 					Operation: to.Ptr("PostClusterCreateScriptActionRequest"),
		// 					ScriptExecutionID: to.Ptr[int64](391144597342127),
		// 					StartTime: to.Ptr("2017-03-22T21:34:39.293"),
		// 					Status: to.Ptr("ValidationFailed"),
		// 			}},
		// 		}
	}
}
