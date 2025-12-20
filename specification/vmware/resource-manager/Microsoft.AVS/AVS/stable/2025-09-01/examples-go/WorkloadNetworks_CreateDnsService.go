package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/WorkloadNetworks_CreateDnsService.json
func ExampleWorkloadNetworksClient_BeginCreateDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadNetworksClient().BeginCreateDNSService(ctx, "group1", "cloud1", "dnsService1", armavs.WorkloadNetworkDNSService{
		Properties: &armavs.WorkloadNetworkDNSServiceProperties{
			DisplayName:    to.Ptr("dnsService1"),
			DNSServiceIP:   to.Ptr("5.5.5.5"),
			DefaultDNSZone: to.Ptr("defaultDnsZone1"),
			FqdnZones: []*string{
				to.Ptr("fqdnZone1"),
			},
			LogLevel: to.Ptr(armavs.DNSServiceLogLevelEnumINFO),
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
	// res = armavs.WorkloadNetworksClientCreateDNSServiceResponse{
	// 	WorkloadNetworkDNSService: &armavs.WorkloadNetworkDNSService{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/workloadNetworks/default/dnsServices/dnsService1"),
	// 		Name: to.Ptr("dnsService1"),
	// 		Properties: &armavs.WorkloadNetworkDNSServiceProperties{
	// 			DisplayName: to.Ptr("dnsService1"),
	// 			DNSServiceIP: to.Ptr("5.5.5.5"),
	// 			DefaultDNSZone: to.Ptr("defaultDnsZone1"),
	// 			FqdnZones: []*string{
	// 				to.Ptr("fqdnZone1"),
	// 			},
	// 			LogLevel: to.Ptr(armavs.DNSServiceLogLevelEnumINFO),
	// 			Status: to.Ptr(armavs.DNSServiceStatusEnumSUCCESS),
	// 			Revision: to.Ptr[int64](1),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/workloadNetworks/dnsServices"),
	// 	},
	// }
}
