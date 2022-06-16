package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/SnapshotsUpdateTags.json
func ExampleSnapshotsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewSnapshotsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.TagsObject{
			Tags: map[string]*string{
				"key2": to.StringPtr("new-val2"),
				"key3": to.StringPtr("val3"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Snapshot.ID: %s\n", *res.ID)
}
