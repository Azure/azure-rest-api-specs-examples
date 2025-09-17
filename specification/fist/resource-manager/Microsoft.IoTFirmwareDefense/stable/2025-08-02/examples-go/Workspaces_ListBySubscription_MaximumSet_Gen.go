package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-08-02/Workspaces_ListBySubscription_MaximumSet_Gen.json
func ExampleWorkspacesClient_NewListBySubscriptionPager_workspacesListBySubscriptionMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListBySubscriptionPager(nil)
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
		// page = armiotfirmwaredefense.WorkspacesClientListBySubscriptionResponse{
		// 	WorkspaceListResult: armiotfirmwaredefense.WorkspaceListResult{
		// 		Value: []*armiotfirmwaredefense.Workspace{
		// 			{
		// 				Location: to.Ptr("East US"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ResourceGroupName/providers/Microsoft.IoTFirmwareDefense/workspaces/WorkspaceName"),
		// 				Properties: &armiotfirmwaredefense.WorkspaceProperties{
		// 					ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armiotfirmwaredefense.SKU{
		// 					Name: to.Ptr("Free"),
		// 					Tier: to.Ptr(armiotfirmwaredefense.SKUTierFree),
		// 					Size: to.Ptr("Free"),
		// 					Family: to.Ptr("Free"),
		// 					Capacity: to.Ptr[int32](20),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key9484": to.Ptr("emizfeemymb"),
		// 				},
		// 				Name: to.Ptr("WorkspaceName"),
		// 				Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces"),
		// 				SystemData: &armiotfirmwaredefense.SystemData{
		// 					CreatedBy: to.Ptr("UserName"),
		// 					CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserName"),
		// 					LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
