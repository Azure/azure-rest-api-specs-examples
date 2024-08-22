package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetEdgeDevices.json
func ExampleEdgeDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEdgeDevicesClient().Get(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeDevicesClientGetResponse{
	// 	                            EdgeDeviceClassification: &armazurestackhci.HciEdgeDevice{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeDevices"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/providers/Microsoft.AzureStackHCI/edgeDevices/default"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 		Kind: to.Ptr(armazurestackhci.DeviceKindHCI),
	// 		Properties: &armazurestackhci.HciEdgeDeviceProperties{
	// 			DeviceConfiguration: &armazurestackhci.DeviceConfiguration{
	// 				DeviceMetadata: to.Ptr(""),
	// 				NicDetails: []*armazurestackhci.NicDetail{
	// 					{
	// 						AdapterName: to.Ptr("ethernet"),
	// 						ComponentID: to.Ptr("VMBUS{f8615163-df3e-46c5-913f-f2d2f965ed0g} "),
	// 						DefaultGateway: to.Ptr("10.10.10.1"),
	// 						DefaultIsolationID: to.Ptr("0"),
	// 						DNSServers: []*string{
	// 							to.Ptr("100.10.10.1")},
	// 							DriverVersion: to.Ptr("10.0.20348.1547 "),
	// 							InterfaceDescription: to.Ptr("NDIS 6.70 "),
	// 							Ip4Address: to.Ptr("10.10.10.10"),
	// 							SubnetMask: to.Ptr("255.255.255.0"),
	// 					}},
	// 				},
	// 				ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			},
	// 		},
	// 		                        }
}
