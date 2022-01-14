Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazuredata%2Farmazuredata%2Fv0.2.0/sdk/resourcemanager/azuredata/armazuredata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// x-ms-original-file: specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/UpdateSqlServerRegistration.json
func ExampleSQLServerRegistrationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazuredata.NewSQLServerRegistrationsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<sql-server-registration-name>",
		armazuredata.SQLServerRegistrationUpdate{
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLServerRegistrationsClientUpdateResult)
}
```
