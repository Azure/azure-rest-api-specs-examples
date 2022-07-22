package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-06-02-preview/examples/Fleets_PatchTags.json
func ExampleFleetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewFleetsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rg1",
		"fleet-1",
		armcontainerservice.FleetPatch{
			Tags: map[string]*string{
				"env":  to.Ptr("prod"),
				"tier": to.Ptr("secure"),
			},
		},
		&armcontainerservice.FleetsClientUpdateOptions{IfMatch: to.Ptr("dfjkwelr7384")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
