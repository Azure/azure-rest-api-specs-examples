package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationsListResult = armcdn.OperationsListResult{
		// 	Value: []*armcdn.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/register/action"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("Registers the Microsoft.Cdn Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Microsoft.Cdn Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/edgenodes/read"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("read"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("EdgeNode"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/edgenodes/write"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("write"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("EdgeNode"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/edgenodes/delete"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("delete"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("EdgeNode"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/profiles/read"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("read"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Profile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/profiles/write"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("write"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Profile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/operationresults/profileresults/write"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("write"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Profile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/operationresults/profileresults/delete"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("delete"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Profile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/operationresults/profileresults/CheckResourceUsage/action"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("CheckResourceUsage"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Profile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cdn/operationresults/profileresults/GenerateSsoUri/action"),
		// 			Display: &armcdn.OperationDisplay{
		// 				Operation: to.Ptr("GenerateSsoUri"),
		// 				Provider: to.Ptr("Microsoft.Cdn"),
		// 				Resource: to.Ptr("Profile"),
		// 			},
		// 	}},
		// }
	}
}
