package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/DeletePeeringService.json
func ExampleServicesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewServicesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<peering-service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
