package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/WorkloadNetworks_UpdateDnsZone.json
func ExampleWorkloadNetworksClient_BeginUpdateDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadNetworksClient().BeginUpdateDNSZone(ctx, "group1", "cloud1", "dnsZone1", armavs.WorkloadNetworkDNSZone{
		Properties: &armavs.WorkloadNetworkDNSZoneProperties{
			DisplayName: to.Ptr("dnsZone1"),
			Domain:      []*string{},
			DNSServerIPs: []*string{
				to.Ptr("1.1.1.1"),
			},
			SourceIP: to.Ptr("8.8.8.8"),
			Revision: to.Ptr[int64](1),
		},
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
	// res = armavs.WorkloadNetworksClientUpdateDNSZoneResponse{
	// 	WorkloadNetworkDNSZone: &armavs.WorkloadNetworkDNSZone{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/dnsZones/dnsZone1"),
	// 		Name: to.Ptr("dnsZone1"),
	// 		Properties: &armavs.WorkloadNetworkDNSZoneProperties{
	// 			DisplayName: to.Ptr("dnsZone1"),
	// 			Domain: []*string{
	// 			},
	// 			DNSServerIPs: []*string{
	// 				to.Ptr("1.1.1.1"),
	// 			},
	// 			SourceIP: to.Ptr("8.8.8.8"),
	// 			DNSServices: to.Ptr[int64](0),
	// 			Revision: to.Ptr[int64](1),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/dnsZones"),
	// 	},
	// }
}
