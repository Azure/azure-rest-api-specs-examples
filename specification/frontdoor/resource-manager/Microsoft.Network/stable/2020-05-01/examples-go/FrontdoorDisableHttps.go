package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorDisableHttps.json
func ExampleFrontendEndpointsClient_BeginDisableHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontendEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDisableHTTPS(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<frontend-endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
