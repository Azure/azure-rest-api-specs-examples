package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceGet.json
func ExampleDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Get(ctx, "rg1", "TestDevice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armhybridnetwork.Device{
	// 	Name: to.Ptr("TestDevice"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/devices"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HybridNetwork/devices/TestDevice"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.AzureStackEdgeFormat{
	// 		DeviceType: to.Ptr(armhybridnetwork.DeviceTypeAzureStackEdge),
	// 		NetworkFunctions: []*armhybridnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.HybridNetwork/networkFunctions/TestNetworkFunction"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		Status: to.Ptr(armhybridnetwork.StatusNotRegistered),
	// 		AzureStackEdge: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg2/providers/Microsoft.DataboxEdge/DataboxEdgeDevices/TestDataboxEdgeDeviceName"),
	// 		},
	// 	},
	// }
}
