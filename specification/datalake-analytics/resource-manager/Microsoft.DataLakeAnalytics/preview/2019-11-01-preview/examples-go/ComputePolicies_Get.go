package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_Get.json
func ExampleComputePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputePoliciesClient().Get(ctx, "contosorg", "contosoadla", "test_policy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComputePolicy = armdatalakeanalytics.ComputePolicy{
	// 	Name: to.Ptr("test_policy"),
	// 	Properties: &armdatalakeanalytics.ComputePolicyProperties{
	// 		MaxDegreeOfParallelismPerJob: to.Ptr[int32](10),
	// 		MinPriorityPerJob: to.Ptr[int32](30),
	// 		ObjectID: to.Ptr("776b9091-8916-4638-87f7-9c989a38da98"),
	// 		ObjectType: to.Ptr(armdatalakeanalytics.AADObjectTypeUser),
	// 	},
	// }
}
