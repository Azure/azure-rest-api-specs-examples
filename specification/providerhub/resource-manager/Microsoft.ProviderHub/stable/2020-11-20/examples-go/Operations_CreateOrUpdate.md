Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv0.2.0/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_CreateOrUpdate.json
func ExampleOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewOperationsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<provider-namespace>",
		armproviderhub.OperationsPutContent{
			Contents: []*armproviderhub.OperationsDefinition{
				{
					Name: to.StringPtr("<name>"),
					Display: &armproviderhub.OperationsDefinitionDisplay{
						Description: to.StringPtr("<description>"),
						Operation:   to.StringPtr("<operation>"),
						Provider:    to.StringPtr("<provider>"),
						Resource:    to.StringPtr("<resource>"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationsClientCreateOrUpdateResult)
}
```
