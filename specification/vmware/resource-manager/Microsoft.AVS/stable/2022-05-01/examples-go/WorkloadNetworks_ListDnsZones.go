package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_ListDnsZones.json
func ExampleWorkloadNetworksClient_NewListDNSZonesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadNetworksClient().NewListDNSZonesPager("group1", "cloud1", nil)
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
		// page.WorkloadNetworkDNSZonesList = armavs.WorkloadNetworkDNSZonesList{
		// 	Value: []*armavs.WorkloadNetworkDNSZone{
		// 		{
		// 			Name: to.Ptr("portMirroring1"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/dnsZones"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/dnsZones/dnsZone1"),
		// 			Properties: &armavs.WorkloadNetworkDNSZoneProperties{
		// 				DisplayName: to.Ptr("dnsZone1"),
		// 				DNSServerIPs: []*string{
		// 					to.Ptr("1.1.1.1")},
		// 					DNSServices: to.Ptr[int64](0),
		// 					Domain: []*string{
		// 					},
		// 					Revision: to.Ptr[int64](1),
		// 					SourceIP: to.Ptr("8.8.8.8"),
		// 				},
		// 		}},
		// 	}
	}
}
