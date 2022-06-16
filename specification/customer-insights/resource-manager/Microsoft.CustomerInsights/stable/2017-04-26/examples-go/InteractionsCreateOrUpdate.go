package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsCreateOrUpdate.json
func ExampleInteractionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewInteractionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestHubRG",
		"sdkTestHub",
		"TestProfileType396",
		armcustomerinsights.InteractionResourceFormat{
			Properties: &armcustomerinsights.InteractionTypeDefinition{
				LargeImage:       to.Ptr("\\\\Images\\\\LargeImage"),
				MediumImage:      to.Ptr("\\\\Images\\\\MediumImage"),
				SmallImage:       to.Ptr("\\\\Images\\\\smallImage"),
				APIEntitySetName: to.Ptr("TestInteractionType6358"),
				Fields: []*armcustomerinsights.PropertyDefinition{
					{
						FieldName:  to.Ptr("TestInteractionType6358"),
						FieldType:  to.Ptr("Edm.String"),
						IsArray:    to.Ptr(false),
						IsRequired: to.Ptr(true),
					},
					{
						FieldName: to.Ptr("profile1"),
						FieldType: to.Ptr("Edm.String"),
					}},
				IDPropertyNames: []*string{
					to.Ptr("TestInteractionType6358")},
				PrimaryParticipantProfilePropertyName: to.Ptr("profile1"),
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
