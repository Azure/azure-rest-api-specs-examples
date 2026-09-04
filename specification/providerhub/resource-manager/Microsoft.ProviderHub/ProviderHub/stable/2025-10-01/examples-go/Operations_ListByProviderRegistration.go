package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/Operations_ListByProviderRegistration.json
func ExampleOperationsClient_ListByProviderRegistration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().ListByProviderRegistration(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.OperationsClientListByProviderRegistrationResponse{
	// 	OperationsPutContent: armproviderhub.OperationsPutContent{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/operations/default"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/operations"),
	// 		Name: to.Ptr("default"),
	// 		Properties: &armproviderhub.OperationsPutContentProperties{
	// 			Contents: []*armproviderhub.LocalizedOperationDefinition{
	// 				{
	// 					Name: to.Ptr("Microsoft.Contoso/Employees/Read"),
	// 					IsDataAction: to.Ptr(false),
	// 					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 							Provider: to.Ptr("Microsoft.Contoso"),
	// 							Resource: to.Ptr("Employees"),
	// 							Operation: to.Ptr("Gets/List employee resources"),
	// 							Description: to.Ptr("Read employees"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Microsoft.Contoso/Employees/Write"),
	// 					IsDataAction: to.Ptr(false),
	// 					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 							Provider: to.Ptr("Microsoft.Contoso"),
	// 							Resource: to.Ptr("Employees"),
	// 							Operation: to.Ptr("Create/update employee resources"),
	// 							Description: to.Ptr("Writes employees"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Microsoft.Contoso/Employees/Delete"),
	// 					IsDataAction: to.Ptr(false),
	// 					Origin: to.Ptr(armproviderhub.OperationOriginsUser),
	// 					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 							Provider: to.Ptr("Microsoft.Contoso"),
	// 							Resource: to.Ptr("Employees"),
	// 							Operation: to.Ptr("Deletes employee resource"),
	// 							Description: to.Ptr("Deletes employees"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Microsoft.Contoso/Employees/Action"),
	// 					IsDataAction: to.Ptr(true),
	// 					Origin: to.Ptr(armproviderhub.OperationOriginsSystem),
	// 					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 							Provider: to.Ptr("Microsoft.Contoso"),
	// 							Resource: to.Ptr("Employees"),
	// 							Operation: to.Ptr("Create/update employee resources"),
	// 							Description: to.Ptr("Writes employees"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
