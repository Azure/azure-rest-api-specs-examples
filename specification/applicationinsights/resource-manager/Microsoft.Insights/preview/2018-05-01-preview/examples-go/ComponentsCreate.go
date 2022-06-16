package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsCreate.json
func ExampleComponentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewComponentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.Component{
			Location: to.StringPtr("<location>"),
			Kind:     to.StringPtr("<kind>"),
			Properties: &armapplicationinsights.ComponentProperties{
				ApplicationType: armapplicationinsights.ApplicationType("web").ToPtr(),
				FlowType:        armapplicationinsights.FlowType("Bluefield").ToPtr(),
				RequestSource:   armapplicationinsights.RequestSource("rest").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ComponentsClientCreateOrUpdateResult)
}
