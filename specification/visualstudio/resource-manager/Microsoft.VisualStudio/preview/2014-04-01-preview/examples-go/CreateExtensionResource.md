Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvisualstudio%2Farmvisualstudio%2Fv0.1.0/sdk/resourcemanager/visualstudio/armvisualstudio/README.md) on how to add the SDK to your project and authenticate.

```go
package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/CreateExtensionResource.json
func ExampleExtensionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvisualstudio.NewExtensionsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-resource-name>",
		"<extension-resource-name>",
		armvisualstudio.ExtensionResourceRequest{
			Location: to.StringPtr("<location>"),
			Plan: &armvisualstudio.ExtensionResourcePlan{
				Name:          to.StringPtr("<name>"),
				Product:       to.StringPtr("<product>"),
				PromotionCode: to.StringPtr("<promotion-code>"),
				Publisher:     to.StringPtr("<publisher>"),
				Version:       to.StringPtr("<version>"),
			},
			Properties: map[string]*string{},
			Tags:       map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ExtensionResource.ID: %s\n", *res.ID)
}
```
