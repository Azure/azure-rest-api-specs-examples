package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listSKUsForExisting.json
func ExampleCapacitiesClient_ListSKUsForCapacity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbidedicated.NewCapacitiesClient("<subscription-id>", cred, nil)
	_, err = client.ListSKUsForCapacity(ctx,
		"<resource-group-name>",
		"<dedicated-capacity-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
