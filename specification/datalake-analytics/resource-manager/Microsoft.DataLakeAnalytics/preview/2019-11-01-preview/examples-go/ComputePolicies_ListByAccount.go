package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_ListByAccount.json
func ExampleComputePoliciesClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewComputePoliciesClient().NewListByAccountPager("contosorg", "contosoadla", nil)
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
		// page.ComputePolicyListResult = armdatalakeanalytics.ComputePolicyListResult{
		// 	Value: []*armdatalakeanalytics.ComputePolicy{
		// 		{
		// 			Name: to.Ptr("test_policy"),
		// 			Properties: &armdatalakeanalytics.ComputePolicyProperties{
		// 				MaxDegreeOfParallelismPerJob: to.Ptr[int32](10),
		// 				MinPriorityPerJob: to.Ptr[int32](30),
		// 				ObjectID: to.Ptr("776b9091-8916-4638-87f7-9c989a38da98"),
		// 				ObjectType: to.Ptr(armdatalakeanalytics.AADObjectTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test_policy1"),
		// 			Properties: &armdatalakeanalytics.ComputePolicyProperties{
		// 				MaxDegreeOfParallelismPerJob: to.Ptr[int32](5),
		// 				MinPriorityPerJob: to.Ptr[int32](15),
		// 				ObjectID: to.Ptr("776b9091-8916-4638-87f7-9c989a38da99"),
		// 				ObjectType: to.Ptr(armdatalakeanalytics.AADObjectTypeGroup),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test_policy2"),
		// 			Properties: &armdatalakeanalytics.ComputePolicyProperties{
		// 				MaxDegreeOfParallelismPerJob: to.Ptr[int32](20),
		// 				MinPriorityPerJob: to.Ptr[int32](60),
		// 				ObjectID: to.Ptr("776b9091-8916-4638-87f7-9c989a38da97"),
		// 				ObjectType: to.Ptr(armdatalakeanalytics.AADObjectTypeServicePrincipal),
		// 			},
		// 	}},
		// }
	}
}
