package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"
)

// x-ms-original-file: specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateLinkServices_ListBySubscriptionId.json
func ExamplePrivateLinkServicesForPowerBIClient_ListBySubscriptionID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbiprivatelinks.NewPrivateLinkServicesForPowerBIClient("<subscription-id>", cred, nil)
	_, err = client.ListBySubscriptionID(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
