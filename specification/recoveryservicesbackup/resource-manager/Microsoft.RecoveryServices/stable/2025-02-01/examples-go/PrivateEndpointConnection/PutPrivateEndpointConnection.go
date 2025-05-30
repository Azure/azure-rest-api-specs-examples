package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionClient().BeginPut(ctx, "gaallavaultbvtd2msi", "gaallaRG", "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b", armrecoveryservicesbackup.PrivateEndpointConnectionResource{
		Properties: &armrecoveryservicesbackup.PrivateEndpointConnection{
			GroupIDs: []*armrecoveryservicesbackup.VaultSubResourceType{
				to.Ptr(armrecoveryservicesbackup.VaultSubResourceTypeAzureBackupSecondary)},
			PrivateEndpoint: &armrecoveryservicesbackup.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"),
			},
			PrivateLinkServiceConnectionState: &armrecoveryservicesbackup.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@company.com"),
				Status:      to.Ptr(armrecoveryservicesbackup.PrivateEndpointConnectionStatusApproved),
			},
			ProvisioningState: to.Ptr(armrecoveryservicesbackup.ProvisioningStateSucceeded),
		},
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
	// res.PrivateEndpointConnectionResource = armrecoveryservicesbackup.PrivateEndpointConnectionResource{
	// 	Name: to.Ptr("gaallatestpe1.3592346090307038890.backup.5704c932-249a-490b-a142-1396838cd3b"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.RecoveryServicesBVTD2/vaults/gaallavaultbvtd2msi/privateEndpointConnections/gaallatestpe3.3592346090307038890.backup.5704c932-249a-490b-a142-1396838cd3b"),
	// 	Properties: &armrecoveryservicesbackup.PrivateEndpointConnection{
	// 		GroupIDs: []*armrecoveryservicesbackup.VaultSubResourceType{
	// 			to.Ptr(armrecoveryservicesbackup.VaultSubResourceTypeAzureBackupSecondary)},
	// 			PrivateEndpoint: &armrecoveryservicesbackup.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armrecoveryservicesbackup.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Approved by johndoe@company.com"),
	// 				Status: to.Ptr(armrecoveryservicesbackup.PrivateEndpointConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armrecoveryservicesbackup.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
