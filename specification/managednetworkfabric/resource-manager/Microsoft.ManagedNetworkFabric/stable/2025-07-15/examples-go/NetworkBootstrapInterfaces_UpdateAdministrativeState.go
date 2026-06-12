package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkBootstrapInterfaces_UpdateAdministrativeState.json
func ExampleNetworkBootstrapInterfacesClient_BeginUpdateAdministrativeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkBootstrapInterfacesClient().BeginUpdateAdministrativeState(ctx, "example-rg", "example-device", "example-interface", armmanagednetworkfabric.UpdateAdministrativeState{
		ResourceIDs: []*string{
			to.Ptr(""),
		},
		State: to.Ptr(armmanagednetworkfabric.EnableDisableStateEnable),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmanagednetworkfabric.NetworkBootstrapInterfacesClientUpdateAdministrativeStateResponse{
	// 	CommonPostActionResponseForStateUpdate: &armmanagednetworkfabric.CommonPostActionResponseForStateUpdate{
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("400"),
	// 			Message: to.Ptr("The request is invalid."),
	// 			Target: to.Ptr("target"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("https://schema.management.azure.com/providers/Microsoft.Management/locations/global/schemas/validationError"),
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 	},
	// }
}
