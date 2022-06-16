package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreateOrUpdateConnectionMonitorTest.json
func ExampleConnectionMonitorTestsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewConnectionMonitorTestsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-service-name>",
		"<connection-monitor-test-name>",
		armpeering.ConnectionMonitorTest{
			Properties: &armpeering.ConnectionMonitorTestProperties{
				Destination:        to.StringPtr("<destination>"),
				DestinationPort:    to.Int32Ptr(443),
				SourceAgent:        to.StringPtr("<source-agent>"),
				TestFrequencyInSec: to.Int32Ptr(30),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConnectionMonitorTestsClientCreateOrUpdateResult)
}
