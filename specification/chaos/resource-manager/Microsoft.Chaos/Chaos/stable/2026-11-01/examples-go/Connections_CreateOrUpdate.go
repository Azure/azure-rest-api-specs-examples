package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Connections_CreateOrUpdate.json
func ExampleConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().CreateOrUpdate(ctx, "exampleRG", "exampleWorkspace", "aksClusterConnection", armchaos.Connection{
		Properties: &armchaos.ConnectionProperties{
			Kind:             to.Ptr(armchaos.ConnectionKindAksExtension),
			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.ContainerService/managedClusters/exampleCluster"),
			PrincipalID:      to.Ptr("00000000-0000-0000-0000-000000000000"),
			TenantID:         to.Ptr("00000000-0000-0000-0000-000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.ConnectionsClientCreateOrUpdateResponse{
	// 	Connection: armchaos.Connection{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/connections/aksClusterConnection"),
	// 		Name: to.Ptr("aksClusterConnection"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/connections"),
	// 		Properties: &armchaos.ConnectionProperties{
	// 			Kind: to.Ptr(armchaos.ConnectionKindAksExtension),
	// 			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.ContainerService/managedClusters/exampleCluster"),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			DataPlaneEndpoint: to.Ptr("https://eastus.dp.chaos.azure.com"),
	// 			Status: to.Ptr(armchaos.ConnectionStatusConnected),
	// 			ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
