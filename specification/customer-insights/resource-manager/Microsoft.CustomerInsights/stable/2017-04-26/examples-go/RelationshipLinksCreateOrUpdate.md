```go
package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksCreateOrUpdate.json
func ExampleRelationshipLinksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewRelationshipLinksClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestHubRG",
		"sdkTestHub",
		"Somelink",
		armcustomerinsights.RelationshipLinkResourceFormat{
			Properties: &armcustomerinsights.RelationshipLinkDefinition{
				Description: map[string]*string{
					"en-us": to.Ptr("Link Description"),
				},
				DisplayName: map[string]*string{
					"en-us": to.Ptr("Link DisplayName"),
				},
				InteractionType: to.Ptr("testInteraction4332"),
				LinkName:        to.Ptr("Somelink"),
				ProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
					{
						InteractionPropertyName: to.Ptr("profile1"),
						ProfilePropertyName:     to.Ptr("ProfileId"),
					}},
				RelatedProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
					{
						InteractionPropertyName: to.Ptr("profile1"),
						ProfilePropertyName:     to.Ptr("ProfileId"),
					}},
				RelationshipName: to.Ptr("testProfile2326994"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomerinsights%2Farmcustomerinsights%2Fv1.0.0/sdk/resourcemanager/customerinsights/armcustomerinsights/README.md) on how to add the SDK to your project and authenticate.
