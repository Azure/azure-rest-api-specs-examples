package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/NetworkSettingsGet.json
func ExampleDevicesClient_GetNetworkSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().GetNetworkSettings(ctx, "testedgedevice", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSettings = armdataboxedge.NetworkSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/networkSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/networkSettings/default"),
	// 	Properties: &armdataboxedge.NetworkSettingsProperties{
	// 		NetworkAdapters: []*armdataboxedge.NetworkAdapter{
	// 			{
	// 				AdapterID: to.Ptr("{47D0D0EC-AA8A-4221-AA2A-355B58082BA5}"),
	// 				AdapterPosition: &armdataboxedge.NetworkAdapterPosition{
	// 					NetworkGroup: to.Ptr(armdataboxedge.NetworkGroupNonRDMA),
	// 					Port: to.Ptr[int32](0),
	// 				},
	// 				DhcpStatus: to.Ptr(armdataboxedge.NetworkAdapterDHCPStatusDisabled),
	// 				DNSServers: []*string{
	// 					to.Ptr("10.50.50.50"),
	// 					to.Ptr("10.50.10.50")},
	// 					Index: to.Ptr[int32](1),
	// 					IPv4Configuration: &armdataboxedge.IPv4Config{
	// 						Gateway: to.Ptr("10.150.76.1"),
	// 						IPAddress: to.Ptr("10.150.78.56"),
	// 						Subnet: to.Ptr("255.255.252.0"),
	// 					},
	// 					IPv6Configuration: &armdataboxedge.IPv6Config{
	// 						Gateway: to.Ptr("fe80::12f3:11ff:fe36:994b%5"),
	// 						IPAddress: to.Ptr("2404:f801:4800:1e:d5c6:50a1:465b:1bbf"),
	// 						PrefixLength: to.Ptr[int32](64),
	// 					},
	// 					IPv6LinkLocalAddress: to.Ptr("fe80::d5c6:50a1:465b:1bbf%5"),
	// 					Label: to.Ptr("DATA1"),
	// 					LinkSpeed: to.Ptr[int64](10000000000),
	// 					MacAddress: to.Ptr("00155D4E265B"),
	// 					NetworkAdapterName: to.Ptr("DATA1"),
	// 					NodeID: to.Ptr("3fd54d9e-f7a0-45bf-bdf1-39b0977d1984"),
	// 					RdmaStatus: to.Ptr(armdataboxedge.NetworkAdapterRDMAStatusIncapable),
	// 					Status: to.Ptr(armdataboxedge.NetworkAdapterStatusInactive),
	// 			}},
	// 		},
	// 	}
}
