package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsUpdate.json
func ExampleDataCollectionEndpointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<data-collection-endpoint-name>",
		&armmonitor.DataCollectionEndpointsClientUpdateOptions{Body: &armmonitor.ResourceForUpdate{
			Tags: map[string]*string{
				"tag1": to.StringPtr("A"),
				"tag2": to.StringPtr("B"),
				"tag3": to.StringPtr("C"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionEndpointsClientUpdateResult)
}
