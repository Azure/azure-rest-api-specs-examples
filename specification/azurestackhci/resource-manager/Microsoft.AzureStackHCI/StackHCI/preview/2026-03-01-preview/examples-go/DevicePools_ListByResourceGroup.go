package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/DevicePools_ListByResourceGroup.json
func ExampleDevicePoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicePoolsClient().NewListByResourceGroupPager("test-rg", nil)
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
		// page = armazurestackhci.DevicePoolsClientListByResourceGroupResponse{
		// 	DevicePoolListResult: armazurestackhci.DevicePoolListResult{
		// 		Value: []*armazurestackhci.DevicePool{
		// 			{
		// 				Properties: &armazurestackhci.DevicePoolProperties{
		// 					ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 					Devices: []*armazurestackhci.DeviceDetail{
		// 						{
		// 							DeviceResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1"),
		// 							ClaimedBy: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/clusters/cluster-1"),
		// 						},
		// 					},
		// 					CustomLocationResourceID: to.Ptr("/subscriptions/59341aac-5917-4ae1-90b3-75f41394d9e2/resourceGroups/BridgeRunner-eus/providers/Microsoft.ExtendedLocation/customLocations/bridgerunnercl-eus"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/devicePools/devicePool-1"),
		// 				Name: to.Ptr("devicePool-1"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/devicePools"),
		// 				SystemData: &armazurestackhci.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T10:46:55.167Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T10:46:55.167Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
