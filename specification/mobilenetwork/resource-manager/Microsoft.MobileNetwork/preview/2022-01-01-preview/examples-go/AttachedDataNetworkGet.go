package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/AttachedDataNetworkGet.json
func ExampleAttachedDataNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewAttachedDataNetworksClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		"<packet-core-data-plane-name>",
		"<attached-data-network-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttachedDataNetworksClientGetResult)
}
