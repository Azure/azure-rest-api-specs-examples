package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsCreate.json
func ExampleDataCollectionEndpointsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<data-collection-endpoint-name>",
		&armmonitor.DataCollectionEndpointsClientCreateOptions{Body: &armmonitor.DataCollectionEndpointResource{
			Location: to.StringPtr("<location>"),
			Properties: &armmonitor.DataCollectionEndpointResourceProperties{
				NetworkACLs: &armmonitor.DataCollectionEndpointNetworkACLs{
					PublicNetworkAccess: armmonitor.KnownPublicNetworkAccessOptions("Enabled").ToPtr(),
				},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionEndpointsClientCreateResult)
}
