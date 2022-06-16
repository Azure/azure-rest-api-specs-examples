package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsUpdateTagsOnly.json
func ExampleComponentsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewComponentsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.TagsResource{
			Tags: map[string]*string{
				"ApplicationGatewayType": to.StringPtr("Internal-Only"),
				"BillingEntity":          to.StringPtr("Self"),
				"Color":                  to.StringPtr("AzureBlue"),
				"CustomField_01":         to.StringPtr("Custom text in some random field named randomly"),
				"NodeType":               to.StringPtr("Edge"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ComponentsClientUpdateTagsResult)
}
