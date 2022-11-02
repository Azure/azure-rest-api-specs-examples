package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_UpdateDnsServices.json
func ExampleWorkloadNetworksClient_BeginUpdateDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewWorkloadNetworksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateDNSService(ctx, "group1", "cloud1", "dnsService1", armavs.WorkloadNetworkDNSService{
		Properties: &armavs.WorkloadNetworkDNSServiceProperties{
			DefaultDNSZone: to.Ptr("defaultDnsZone1"),
			DisplayName:    to.Ptr("dnsService1"),
			DNSServiceIP:   to.Ptr("5.5.5.5"),
			FqdnZones: []*string{
				to.Ptr("fqdnZone1")},
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
	// TODO: use response item
	_ = res
}
