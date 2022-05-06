Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazuredata%2Farmazuredata%2Fv0.4.0/sdk/resourcemanager/azuredata/armazuredata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/CreateOrUpdateSqlServerWithRegistrationGroup.json
func ExampleSQLServersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazuredata.NewSQLServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-server-registration-name>",
		"<sql-server-name>",
		armazuredata.SQLServer{
			Properties: &armazuredata.SQLServerProperties{
				Cores:          to.Ptr[int32](8),
				Edition:        to.Ptr("<edition>"),
				PropertyBag:    to.Ptr("<property-bag>"),
				RegistrationID: to.Ptr("<registration-id>"),
				Version:        to.Ptr("<version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
