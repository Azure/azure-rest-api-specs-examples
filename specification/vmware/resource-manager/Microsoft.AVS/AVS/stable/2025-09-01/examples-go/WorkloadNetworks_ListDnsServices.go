package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/WorkloadNetworks_ListDnsServices.json
func ExampleWorkloadNetworksClient_NewListDNSServicesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadNetworksClient().NewListDNSServicesPager("group1", "cloud1", nil)
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
		// page = armavs.WorkloadNetworksClientListDNSServicesResponse{
		// 	WorkloadNetworkDNSServicesList: armavs.WorkloadNetworkDNSServicesList{
		// 		Value: []*armavs.WorkloadNetworkDNSService{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/dnsServices/dnsService1"),
		// 				Name: to.Ptr("portMirroring1"),
		// 				Properties: &armavs.WorkloadNetworkDNSServiceProperties{
		// 					DisplayName: to.Ptr("dnsService1"),
		// 					DNSServiceIP: to.Ptr("5.5.5.5"),
		// 					DefaultDNSZone: to.Ptr("defaultDnsZone1"),
		// 					FqdnZones: []*string{
		// 						to.Ptr("fqdnZone1"),
		// 					},
		// 					LogLevel: to.Ptr(armavs.DNSServiceLogLevelEnumINFO),
		// 					Status: to.Ptr(armavs.DNSServiceStatusEnumSUCCESS),
		// 					Revision: to.Ptr[int64](1),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/dnsServices"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
