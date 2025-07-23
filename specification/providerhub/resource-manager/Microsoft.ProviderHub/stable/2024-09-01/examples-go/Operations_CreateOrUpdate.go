package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/Operations_CreateOrUpdate.json
func ExampleOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().CreateOrUpdate(ctx, "Microsoft.Contoso", armproviderhub.OperationsPutContent{
		Properties: &armproviderhub.OperationsPutContentProperties{
			Contents: []*armproviderhub.LocalizedOperationDefinition{
				{
					Name:       to.Ptr("RP.69C09791/register/action"),
					ActionType: to.Ptr(armproviderhub.OperationActionTypeInternal),
					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
							Description: to.Ptr("Registers the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
							Operation:   to.Ptr("Registers the RP.69C09791 Resource Provider"),
							Provider:    to.Ptr("RP.69C09791"),
							Resource:    to.Ptr("Register"),
						},
					},
					IsDataAction: to.Ptr(true),
				},
				{
					Name: to.Ptr("RP.69C09791/unregister/action"),
					Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
						Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
							Description: to.Ptr("Unregisters the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
							Operation:   to.Ptr("Unregisters the RP.69C09791 Resource Provider"),
							Provider:    to.Ptr("RP.69C09791"),
							Resource:    to.Ptr("Unregister"),
						},
						En: &armproviderhub.LocalizedOperationDisplayDefinitionEn{
							Description: to.Ptr("ece249f5-b5b9-492d-ac68-b4e1be1677bc"),
							Operation:   to.Ptr("d31623d6-8765-42fb-aca2-5a58303e52dd"),
							Provider:    to.Ptr("RP.69C09791"),
							Resource:    to.Ptr("2e1803d4-417f-492c-b305-148da38b211e"),
						},
					},
					IsDataAction: to.Ptr(false),
					Origin:       to.Ptr(armproviderhub.OperationOriginsSystem),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationsPutContent = armproviderhub.OperationsPutContent{
	// 	Name: to.Ptr("operationTest"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/operations"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/operations/default"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.OperationsPutContentProperties{
	// 		Contents: []*armproviderhub.LocalizedOperationDefinition{
	// 			{
	// 				Name: to.Ptr("RP.69C09791/register/action"),
	// 				ActionType: to.Ptr(armproviderhub.OperationActionTypeInternal),
	// 				Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 					Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 						Description: to.Ptr("Registers the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
	// 						Operation: to.Ptr("Registers the RP.69C09791 Resource Provider"),
	// 						Provider: to.Ptr("RP.69C09791"),
	// 						Resource: to.Ptr("Register"),
	// 					},
	// 				},
	// 				IsDataAction: to.Ptr(true),
	// 			},
	// 			{
	// 				Name: to.Ptr("RP.69C09791/unregister/action"),
	// 				Display: &armproviderhub.LocalizedOperationDefinitionDisplay{
	// 					Default: &armproviderhub.LocalizedOperationDisplayDefinitionDefault{
	// 						Description: to.Ptr("Unregisters the subscription for the RP.69C09791 resource provider and enables the creation of RP.69C09791."),
	// 						Operation: to.Ptr("Unregisters the RP.69C09791 Resource Provider"),
	// 						Provider: to.Ptr("RP.69C09791"),
	// 						Resource: to.Ptr("Unregister"),
	// 					},
	// 					En: &armproviderhub.LocalizedOperationDisplayDefinitionEn{
	// 						Description: to.Ptr("ece249f5-b5b9-492d-ac68-b4e1be1677bc"),
	// 						Operation: to.Ptr("d31623d6-8765-42fb-aca2-5a58303e52dd"),
	// 						Provider: to.Ptr("RP.69C09791"),
	// 						Resource: to.Ptr("2e1803d4-417f-492c-b305-148da38b211e"),
	// 					},
	// 				},
	// 				IsDataAction: to.Ptr(false),
	// 				Origin: to.Ptr(armproviderhub.OperationOriginsSystem),
	// 		}},
	// 	},
	// }
}
