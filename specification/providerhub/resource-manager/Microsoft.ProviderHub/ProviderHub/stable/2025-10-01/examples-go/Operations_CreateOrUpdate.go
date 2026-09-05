package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/Operations_CreateOrUpdate.json
func ExampleOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().CreateOrUpdate(ctx, "Microsoft.Contoso", armproviderhub.OperationsPutContent{
		Properties: &armproviderhub.OperationsPutContentProperties{
			Contents: []*armproviderhub.LocalizedOperationDefinition{
				{
					Name:         to.Ptr("RP.69C09791/register/action"),
					IsDataAction: to.Ptr(true),
					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
							Provider:    to.Ptr("RP.69C09791"),
							Resource:    to.Ptr("Register"),
							Operation:   to.Ptr("Registers the RP.69C09791 Resource Provider"),
							Description: to.Ptr("Registers the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
						},
					},
					ActionType: to.Ptr(armproviderhub.OperationActionTypeInternal),
				},
				{
					Name:         to.Ptr("RP.69C09791/unregister/action"),
					IsDataAction: to.Ptr(false),
					Origin:       to.Ptr(armproviderhub.OperationOriginsSystem),
					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
							Provider:    to.Ptr("RP.69C09791"),
							Resource:    to.Ptr("Unregister"),
							Operation:   to.Ptr("Unregisters the RP.69C09791 Resource Provider"),
							Description: to.Ptr("Unregisters the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
						},
						En: &armproviderhub.LocalizedOperationDisplayDefinitionEn{
							Provider:    to.Ptr("RP.69C09791"),
							Resource:    to.Ptr("2e1803d4-417f-492c-b305-148da38b211e"),
							Operation:   to.Ptr("d31623d6-8765-42fb-aca2-5a58303e52dd"),
							Description: to.Ptr("ece249f5-b5b9-492d-ac68-b4e1be1677bc"),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.OperationsClientCreateOrUpdateResponse{
	// 	OperationsPutContent: armproviderhub.OperationsPutContent{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/operations/default"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/operations"),
	// 		Name: to.Ptr("operationTest"),
	// 		Properties: &armproviderhub.OperationsPutContentProperties{
	// 			Contents: []*armproviderhub.LocalizedOperationDefinition{
	// 				{
	// 					Name: to.Ptr("RP.69C09791/register/action"),
	// 					IsDataAction: to.Ptr(true),
	// 					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 							Provider: to.Ptr("RP.69C09791"),
	// 							Resource: to.Ptr("Register"),
	// 							Operation: to.Ptr("Registers the RP.69C09791 Resource Provider"),
	// 							Description: to.Ptr("Registers the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
	// 						},
	// 					},
	// 					ActionType: to.Ptr(armproviderhub.OperationActionTypeInternal),
	// 				},
	// 				{
	// 					Name: to.Ptr("RP.69C09791/unregister/action"),
	// 					IsDataAction: to.Ptr(false),
	// 					Origin: to.Ptr(armproviderhub.OperationOriginsSystem),
	// 					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 							Provider: to.Ptr("RP.69C09791"),
	// 							Resource: to.Ptr("Unregister"),
	// 							Operation: to.Ptr("Unregisters the RP.69C09791 Resource Provider"),
	// 							Description: to.Ptr("Unregisters the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
	// 						},
	// 						En: &armproviderhub.LocalizedOperationDisplayDefinitionEn{
	// 							Provider: to.Ptr("RP.69C09791"),
	// 							Resource: to.Ptr("2e1803d4-417f-492c-b305-148da38b211e"),
	// 							Operation: to.Ptr("d31623d6-8765-42fb-aca2-5a58303e52dd"),
	// 							Description: to.Ptr("ece249f5-b5b9-492d-ac68-b4e1be1677bc"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armproviderhub.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 		},
	// 	},
	// }
}
