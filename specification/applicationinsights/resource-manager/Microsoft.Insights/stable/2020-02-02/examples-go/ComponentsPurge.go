package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsPurge.json
func ExampleComponentsClient_Purge() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewComponentsClient("00000000-0000-0000-0000-00000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Purge(ctx,
		"OIAutoRest5123",
		"aztest5048",
		armapplicationinsights.ComponentPurgeBody{
			Filters: []*armapplicationinsights.ComponentPurgeBodyFilters{
				{
					Column:   to.Ptr("TimeGenerated"),
					Operator: to.Ptr(">"),
					Value:    "2017-09-01T00:00:00",
				}},
			Table: to.Ptr("Heartbeat"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
