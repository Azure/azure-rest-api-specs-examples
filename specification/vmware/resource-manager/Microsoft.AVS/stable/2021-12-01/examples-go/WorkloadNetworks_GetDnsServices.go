package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetDnsServices.json
func ExampleWorkloadNetworksClient_GetDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetDNSService(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-service-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSService.ID: %s\n", *res.ID)
}
