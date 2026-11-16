package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2026-07-01-preview/FabricShortcuts/FabricShortcutCreate.json
func ExampleFabricShortcutsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFabricShortcutsClient().CreateOrUpdate(ctx, "contoso-rg", "contoso-eventhub", "orders", "orders-shortcut", armeventhub.FabricShortcut{
		Properties: &armeventhub.FabricShortcutProperties{
			Configuration: &armeventhub.FabricShortcutConfiguration{
				ArtifactID:        to.Ptr("33333333-3333-3333-3333-333333333333"),
				ArtifactName:      to.Ptr("orders-eventstream"),
				PremiumCapacityID: to.Ptr("44444444-4444-4444-4444-444444444444"),
				TenantID:          to.Ptr("11111111-1111-1111-1111-111111111111"),
				WorkspaceID:       to.Ptr("22222222-2222-2222-2222-222222222222"),
				WorkspaceName:     to.Ptr("contoso-workspace"),
			},
			ShortcutStatus: to.Ptr(armeventhub.FabricShortcutStatusPending),
			ShortcutType:   to.Ptr(armeventhub.FabricShortcutTypeEntity),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventhub.FabricShortcutsClientCreateOrUpdateResponse{
	// 	FabricShortcut: armeventhub.FabricShortcut{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.EventHub/namespaces/contoso-eventhub/eventhubs/orders/fabricShortcuts/orders-shortcut"),
	// 		Location: to.Ptr("southcentralus"),
	// 		Name: to.Ptr("orders-shortcut"),
	// 		Properties: &armeventhub.FabricShortcutProperties{
	// 			Configuration: &armeventhub.FabricShortcutConfiguration{
	// 				ArtifactID: to.Ptr("33333333-3333-3333-3333-333333333333"),
	// 				ArtifactName: to.Ptr("orders-eventstream"),
	// 				PremiumCapacityID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 				TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				WorkspaceID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				WorkspaceName: to.Ptr("contoso-workspace"),
	// 			},
	// 			CreatedAt: to.Ptr(time.Date(2026, time.July, 1, 12, 0, 0, 0, time.UTC)),
	// 			ModifiedAt: to.Ptr(time.Date(2026, time.July, 1, 12, 0, 0, 0, time.UTC)),
	// 			ShortcutStatus: to.Ptr(armeventhub.FabricShortcutStatusPending),
	// 			ShortcutType: to.Ptr(armeventhub.FabricShortcutTypeEntity),
	// 			StatusDescription: to.Ptr("Pending approval"),
	// 		},
	// 		Type: to.Ptr("Microsoft.EventHub/namespaces/eventhubs/fabricShortcuts"),
	// 	},
	// }
}
