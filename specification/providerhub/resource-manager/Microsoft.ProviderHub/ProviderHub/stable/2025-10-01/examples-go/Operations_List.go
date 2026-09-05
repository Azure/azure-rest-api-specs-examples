package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armproviderhub.OperationsClientListResponse{
		// 	OperationsDefinitionArrayResponseWithContinuation: armproviderhub.OperationsDefinitionArrayResponseWithContinuation{
		// 		Value: []*armproviderhub.OperationsDefinition{
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/register/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("register"),
		// 					Operation: to.Ptr("Register for Microsoft.ProviderHub"),
		// 					Description: to.Ptr("Registers the specified subscription with Microsoft.ProviderHub resource provider"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("defaultRollouts"),
		// 					Operation: to.Ptr("Create or Update rollout"),
		// 					Description: to.Ptr("Creates or Updates any rollout"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("defaultRollouts"),
		// 					Operation: to.Ptr("Read rollout"),
		// 					Description: to.Ptr("Reads any rollout"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("defaultRollouts"),
		// 					Operation: to.Ptr("Delete rollout"),
		// 					Description: to.Ptr("Deletes any rollout"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/stop/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("defaultRollouts"),
		// 					Operation: to.Ptr("Delete rollout"),
		// 					Description: to.Ptr("Deletes any rollout"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/customRollouts/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("customRollouts"),
		// 					Operation: to.Ptr("Create or Update rollout"),
		// 					Description: to.Ptr("Creates or Updates any rollout"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ProviderHub/customRollouts/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armproviderhub.OperationsDefinitionDisplay{
		// 					Provider: to.Ptr("Microsoft ProviderHub"),
		// 					Resource: to.Ptr("customRollouts"),
		// 					Operation: to.Ptr("Read rollout"),
		// 					Description: to.Ptr("Reads any rollout"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
