package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/DevicePools_CreateOrUpdate.json
func ExampleDevicePoolsClient_BeginCreateOrUpdate_devicePoolsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDevicePoolsClient().BeginCreateOrUpdate(ctx, "ArcInstance-rg", "devicePool-1", armazurestackhci.DevicePool{
		Properties: &armazurestackhci.DevicePoolProperties{},
		Location:   to.Ptr("eastus"),
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
	// res = armazurestackhci.DevicePoolsClientCreateOrUpdateResponse{
	// 	DevicePool: &armazurestackhci.DevicePool{
	// 		Properties: &armazurestackhci.DevicePoolProperties{
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			Devices: []*armazurestackhci.DeviceDetail{
	// 				{
	// 					DeviceResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1"),
	// 					ClaimedBy: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/clusters/cluster-1"),
	// 				},
	// 			},
	// 			CustomLocationResourceID: to.Ptr("/subscriptions/59341aac-5917-4ae1-90b3-75f41394d9e2/resourceGroups/BridgeRunner-eus/providers/Microsoft.ExtendedLocation/customLocations/bridgerunnercl-eus"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/devicePools/devicePool-1"),
	// 		Name: to.Ptr("devicePool-1"),
	// 		Type: to.Ptr("bwcgsqpmknwznzsscflmtyccoli"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("rxhpfyokywvjgo"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T09:48:11.362Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("g"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T09:48:11.363Z"); return t}()),
	// 		},
	// 	},
	// }
}
