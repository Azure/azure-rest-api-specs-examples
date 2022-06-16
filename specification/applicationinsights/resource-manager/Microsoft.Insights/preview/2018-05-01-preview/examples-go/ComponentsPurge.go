package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsPurge.json
func ExampleComponentsClient_Purge() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewComponentsClient("<subscription-id>", cred, nil)
	_, err = client.Purge(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.ComponentPurgeBody{
			Filters: []*armapplicationinsights.ComponentPurgeBodyFilters{
				{
					Column:   to.StringPtr("<column>"),
					Operator: to.StringPtr("<operator>"),
					Value:    "2017-09-01T00:00:00",
				}},
			Table: to.StringPtr("<table>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
