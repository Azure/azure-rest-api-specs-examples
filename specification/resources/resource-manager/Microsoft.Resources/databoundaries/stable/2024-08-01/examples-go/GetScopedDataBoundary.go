package armdataboundaries_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoundaries/armdataboundaries"
)

// Generated from example definition: 2024-08-01/GetScopedDataBoundary.json
func ExampleClient_GetScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboundaries.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GetScope(ctx, "subscriptions/11111111-1111-1111-1111-111111111111/resourcegroups/my-resource-group", armdataboundaries.DefaultNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataboundaries.ClientGetScopeResponse{
	// 	DataBoundaryDefinition: armdataboundaries.DataBoundaryDefinition{
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ID: to.Ptr("/providers/Microsoft.Resources/dataBoundaries/00000000-0000-0000-0000-000000000000"),
	// 		Properties: &armdataboundaries.DataBoundaryProperties{
	// 			DataBoundary: to.Ptr(armdataboundaries.DataBoundaryEU),
	// 			ProvisioningState: to.Ptr(armdataboundaries.ProvisioningState("Created")),
	// 		},
	// 	},
	// }
}
