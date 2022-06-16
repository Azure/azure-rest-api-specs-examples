package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/LookingGlassInvokeCommand.json
func ExampleLookingGlassClient_Invoke() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewLookingGlassClient("<subscription-id>", cred, nil)
	res, err := client.Invoke(ctx,
		armpeering.LookingGlassCommand("Traceroute"),
		armpeering.LookingGlassSourceType("AzureRegion"),
		"<source-location>",
		"<destination-ip>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LookingGlassClientInvokeResult)
}
