```go
package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/CreateOrUpdateSqlServerRegistration.json
func ExampleSQLServerRegistrationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armazuredata.NewSQLServerRegistrationsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testrg",
		"testsqlregistration",
		armazuredata.SQLServerRegistration{
			Location: to.Ptr("northeurope"),
			Tags: map[string]*string{
				"mytag": to.Ptr("myval"),
			},
			Properties: &armazuredata.SQLServerRegistrationProperties{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazuredata%2Farmazuredata%2Fv0.5.0/sdk/resourcemanager/azuredata/armazuredata/README.md) on how to add the SDK to your project and authenticate.
