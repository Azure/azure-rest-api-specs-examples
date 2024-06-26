package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/listOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armpowerplatform.OperationListResult{
		// 	Value: []*armpowerplatform.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/write"),
		// 			Display: &armpowerplatform.OperationDisplay{
		// 				Description: to.Ptr("Create new enterprisePolicy."),
		// 				Operation: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/write"),
		// 				Provider: to.Ptr("Microsoft PowerPlatform"),
		// 				Resource: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/read"),
		// 			Display: &armpowerplatform.OperationDisplay{
		// 				Description: to.Ptr("Get enterprisePolicy."),
		// 				Operation: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/read"),
		// 				Provider: to.Ptr("Microsoft PowerPlatform"),
		// 				Resource: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerPlatform/accounts/write"),
		// 			Display: &armpowerplatform.OperationDisplay{
		// 				Description: to.Ptr("Create new account."),
		// 				Operation: to.Ptr("Microsoft.PowerPlatform/accounts/write"),
		// 				Provider: to.Ptr("Microsoft PowerPlatform"),
		// 				Resource: to.Ptr("Microsoft.PowerPlatform/accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerPlatform/accounts/read"),
		// 			Display: &armpowerplatform.OperationDisplay{
		// 				Description: to.Ptr("Get account."),
		// 				Operation: to.Ptr("Microsoft.PowerPlatform/accounts/read"),
		// 				Provider: to.Ptr("Microsoft PowerPlatform"),
		// 				Resource: to.Ptr("Microsoft.PowerPlatform/accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
