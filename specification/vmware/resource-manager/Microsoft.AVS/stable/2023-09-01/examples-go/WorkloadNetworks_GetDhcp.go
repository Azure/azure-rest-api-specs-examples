package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_GetDhcp.json
func ExampleWorkloadNetworksClient_GetDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadNetworksClient().GetDhcp(ctx, "group1", "dhcp1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkloadNetworkDhcp = armavs.WorkloadNetworkDhcp{
	// 	Name: to.Ptr("dhcp1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/dhcpConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/dhcpConfigurations/dhcpConfigurations1"),
	// 	Properties: &armavs.WorkloadNetworkDhcpServer{
	// 		DhcpType: to.Ptr(armavs.DhcpTypeEnumSERVER),
	// 		DisplayName: to.Ptr("dhcpConfigurations1"),
	// 		Revision: to.Ptr[int64](1),
	// 		Segments: []*string{
	// 			to.Ptr("segment1"),
	// 			to.Ptr("segment2")},
	// 			LeaseTime: to.Ptr[int64](86400),
	// 			ServerAddress: to.Ptr("40.1.5.1/24"),
	// 		},
	// 	}
}
