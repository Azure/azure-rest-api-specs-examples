package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Connections_ListAll.json
func ExampleConnectionsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListAllPager("exampleRG", "exampleWorkspace", nil)
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
		// page = armchaos.ConnectionsClientListAllResponse{
		// 	ConnectionListResult: armchaos.ConnectionListResult{
		// 		Value: []*armchaos.Connection{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/connections/aksClusterConnection"),
		// 				Name: to.Ptr("aksClusterConnection"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/connections"),
		// 				Properties: &armchaos.ConnectionProperties{
		// 					Kind: to.Ptr(armchaos.ConnectionKindAksExtension),
		// 					TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.ContainerService/managedClusters/exampleCluster"),
		// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					DataPlaneEndpoint: to.Ptr("https://eastus.dp.chaos.azure.com"),
		// 					Status: to.Ptr(armchaos.ConnectionStatusConnected),
		// 					ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/connections/agentConnection"),
		// 				Name: to.Ptr("agentConnection"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/connections"),
		// 				Properties: &armchaos.ConnectionProperties{
		// 					Kind: to.Ptr(armchaos.ConnectionKindChaosAgent),
		// 					TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVm"),
		// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					CertificateSubjectName: to.Ptr("CN=exampleVm.agent.chaos.azure.com"),
		// 					CertificateIssuer: to.Ptr("CN=CACert"),
		// 					Status: to.Ptr(armchaos.ConnectionStatusPending),
		// 					ProvisioningState: to.Ptr(armchaos.ProvisioningStateCreating),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 16, 9, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 16, 9, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
