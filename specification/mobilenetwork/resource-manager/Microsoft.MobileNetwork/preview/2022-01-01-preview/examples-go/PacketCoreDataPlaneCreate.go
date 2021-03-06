package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/PacketCoreDataPlaneCreate.json
func ExamplePacketCoreDataPlanesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewPacketCoreDataPlanesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		"<packet-core-data-plane-name>",
		armmobilenetwork.PacketCoreDataPlane{
			Location: to.StringPtr("<location>"),
			Properties: &armmobilenetwork.PacketCoreDataPlanePropertiesFormat{
				UserPlaneAccessInterface: &armmobilenetwork.InterfaceProperties{
					Name: to.StringPtr("<name>"),
				},
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
	log.Printf("Response result: %#v\n", res.PacketCoreDataPlanesClientCreateOrUpdateResult)
}
