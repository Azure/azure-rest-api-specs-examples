```go
package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfilesCreateOrUpdate.json
func ExampleProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewProfilesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestHubRG",
		"sdkTestHub",
		"TestProfileType396",
		armcustomerinsights.ProfileResourceFormat{
			Properties: &armcustomerinsights.ProfileTypeDefinition{
				LargeImage:       to.Ptr("\\\\Images\\\\LargeImage"),
				MediumImage:      to.Ptr("\\\\Images\\\\MediumImage"),
				SmallImage:       to.Ptr("\\\\Images\\\\smallImage"),
				APIEntitySetName: to.Ptr("TestProfileType396"),
				Fields: []*armcustomerinsights.PropertyDefinition{
					{
						FieldName:  to.Ptr("Id"),
						FieldType:  to.Ptr("Edm.String"),
						IsArray:    to.Ptr(false),
						IsRequired: to.Ptr(true),
					},
					{
						FieldName:  to.Ptr("ProfileId"),
						FieldType:  to.Ptr("Edm.String"),
						IsArray:    to.Ptr(false),
						IsRequired: to.Ptr(true),
					},
					{
						FieldName:  to.Ptr("LastName"),
						FieldType:  to.Ptr("Edm.String"),
						IsArray:    to.Ptr(false),
						IsRequired: to.Ptr(true),
					},
					{
						FieldName:  to.Ptr("TestProfileType396"),
						FieldType:  to.Ptr("Edm.String"),
						IsArray:    to.Ptr(false),
						IsRequired: to.Ptr(true),
					},
					{
						FieldName:  to.Ptr("SavingAccountBalance"),
						FieldType:  to.Ptr("Edm.Int32"),
						IsArray:    to.Ptr(false),
						IsRequired: to.Ptr(true),
					}},
				SchemaItemTypeLink: to.Ptr("SchemaItemTypeLink"),
				StrongIDs: []*armcustomerinsights.StrongID{
					{
						KeyPropertyNames: []*string{
							to.Ptr("Id"),
							to.Ptr("SavingAccountBalance")},
						StrongIDName: to.Ptr("Id"),
					},
					{
						KeyPropertyNames: []*string{
							to.Ptr("ProfileId"),
							to.Ptr("LastName")},
						StrongIDName: to.Ptr("ProfileId"),
					}},
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
