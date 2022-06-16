package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Update.json
func ExampleOrganizationClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfluent.NewOrganizationClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<organization-name>",
		&armconfluent.OrganizationClientUpdateOptions{Body: &armconfluent.OrganizationResourceUpdate{
			Tags: map[string]*string{
				"client": to.StringPtr("dev-client"),
				"env":    to.StringPtr("dev"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OrganizationClientUpdateResult)
}
