package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetScriptActionById.json
func ExampleScriptActionsClient_GetExecutionDetail() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptActionsClient().GetExecutionDetail(ctx, "rg1", "cluster1", "391145124054712", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuntimeScriptActionDetail = armhdinsight.RuntimeScriptActionDetail{
	// 	Name: to.Ptr("Test"),
	// 	ApplicationName: to.Ptr("app1"),
	// 	Roles: []*string{
	// 		to.Ptr("headnode"),
	// 		to.Ptr("workernode")},
	// 		URI: to.Ptr("http://testurl.com/install.ssh"),
	// 		DebugInformation: to.Ptr(""),
	// 		EndTime: to.Ptr("2017-03-22T21:34:39.293"),
	// 		ExecutionSummary: []*armhdinsight.ScriptActionExecutionSummary{
	// 		},
	// 		Operation: to.Ptr("PostClusterCreateScriptActionRequest"),
	// 		ScriptExecutionID: to.Ptr[int64](391145124054712),
	// 		StartTime: to.Ptr("2017-03-22T21:34:39.293"),
	// 		Status: to.Ptr("ValidationFailed"),
	// 	}
}
