package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/LookingGlassInvokeCommand.json
func ExampleLookingGlassClient_Invoke() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpeering.NewLookingGlassClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Invoke(ctx,
		armpeering.LookingGlassCommandTraceroute,
		armpeering.LookingGlassSourceTypeAzureRegion,
		"West US",
		"0.0.0.0",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
