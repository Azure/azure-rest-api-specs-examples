package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationsDefinitionArrayResponseWithContinuation = armproviderhub.OperationsDefinitionArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.OperationsDefinition{
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/register/action"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Registers the specified subscription with Microsoft.ProviderHub resource provider"),
		// 				Operation: to.Ptr("Register for Microsoft.ProviderHub"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("register"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/write"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Creates or Updates any rollout"),
		// 				Operation: to.Ptr("Create or Update rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/read"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Reads any rollout"),
		// 				Operation: to.Ptr("Read rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/delete"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Deletes any rollout"),
		// 				Operation: to.Ptr("Delete rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/stop/action"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Deletes any rollout"),
		// 				Operation: to.Ptr("Delete rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/customRollouts/write"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Creates or Updates any rollout"),
		// 				Operation: to.Ptr("Create or Update rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("customRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/customRollouts/read"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Reads any rollout"),
		// 				Operation: to.Ptr("Read rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("customRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
