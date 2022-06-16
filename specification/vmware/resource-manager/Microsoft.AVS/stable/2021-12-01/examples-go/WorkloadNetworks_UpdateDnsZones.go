package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDnsZones.json
func ExampleWorkloadNetworksClient_BeginUpdateDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDNSZone(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-zone-id>",
		armavs.WorkloadNetworkDNSZone{
			Properties: &armavs.WorkloadNetworkDNSZoneProperties{
				DisplayName: to.StringPtr("<display-name>"),
				DNSServerIPs: []*string{
					to.StringPtr("1.1.1.1")},
				Domain:   []*string{},
				Revision: to.Int64Ptr(1),
				SourceIP: to.StringPtr("<source-ip>"),
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
	log.Printf("WorkloadNetworkDNSZone.ID: %s\n", *res.ID)
}
