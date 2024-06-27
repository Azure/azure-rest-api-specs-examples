package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_GetDnsZone.json
func ExampleWorkloadNetworksClient_GetDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadNetworksClient().GetDNSZone(ctx, "group1", "cloud1", "dnsZone1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkloadNetworkDNSZone = armavs.WorkloadNetworkDNSZone{
	// 	Name: to.Ptr("portMirroring1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/dnsZones"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/dnsZones/dnsZone1"),
	// 	Properties: &armavs.WorkloadNetworkDNSZoneProperties{
	// 		DisplayName: to.Ptr("dnsZone1"),
	// 		DNSServerIPs: []*string{
	// 			to.Ptr("1.1.1.1")},
	// 			DNSServices: to.Ptr[int64](0),
	// 			Domain: []*string{
	// 			},
	// 			Revision: to.Ptr[int64](1),
	// 			SourceIP: to.Ptr("8.8.8.8"),
	// 		},
	// 	}
}
