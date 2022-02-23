Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.1/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSchemas_CreateOrUpdate.json
func ExampleIntegrationAccountSchemasClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountSchemasClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<schema-name>",
		armlogic.IntegrationAccountSchema{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"integrationAccountSchemaName": to.StringPtr("IntegrationAccountSchema8120"),
			},
			Properties: &armlogic.IntegrationAccountSchemaProperties{
				Content:     to.StringPtr("<content>"),
				ContentType: to.StringPtr("<content-type>"),
				Metadata:    map[string]interface{}{},
				SchemaType:  armlogic.SchemaType("Xml").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationAccountSchemasClientCreateOrUpdateResult)
}
```
