Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvisualstudio%2Farmvisualstudio%2Fv0.2.1/sdk/resourcemanager/visualstudio/armvisualstudio/README.md) on how to add the SDK to your project and authenticate.

```go
package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/CreateResource.json
func ExampleAccountsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvisualstudio.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armvisualstudio.AccountResourceRequest{
			AccountName:   to.StringPtr("<account-name>"),
			Location:      to.StringPtr("<location>"),
			OperationType: armvisualstudio.AccountResourceRequestOperationTypeCreate.ToPtr(),
			Properties:    map[string]*string{},
			Tags:          map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientCreateOrUpdateResult)
}
```
