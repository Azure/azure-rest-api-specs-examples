package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDnsServices.json
func ExampleWorkloadNetworksClient_BeginUpdateDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDNSService(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-service-id>",
		armavs.WorkloadNetworkDNSService{
			Properties: &armavs.WorkloadNetworkDNSServiceProperties{
				DefaultDNSZone: to.StringPtr("<default-dnszone>"),
				DisplayName:    to.StringPtr("<display-name>"),
				DNSServiceIP:   to.StringPtr("<dnsservice-ip>"),
				FqdnZones: []*string{
					to.StringPtr("fqdnZone1")},
				LogLevel: armavs.DNSServiceLogLevelEnumINFO.ToPtr(),
				Revision: to.Int64Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSService.ID: %s\n", *res.ID)
}
