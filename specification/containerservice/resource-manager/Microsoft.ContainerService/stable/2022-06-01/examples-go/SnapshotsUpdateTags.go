package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-06-01/examples/SnapshotsUpdateTags.json
func ExampleSnapshotsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewSnapshotsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx,
		"rg1",
		"snapshot1",
		armcontainerservice.TagsObject{
			Tags: map[string]*string{
				"key2": to.Ptr("new-val2"),
				"key3": to.Ptr("val3"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
