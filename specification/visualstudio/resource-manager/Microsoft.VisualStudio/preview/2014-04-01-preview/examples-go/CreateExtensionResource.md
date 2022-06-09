```go
package armvisualstudio_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/CreateExtensionResource.json
func ExampleExtensionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvisualstudio.NewExtensionsClient("0de7f055-dbea-498d-8e9e-da287eedca90", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"VS-Example-Group",
		"ExampleAccount",
		"ms.example",
		armvisualstudio.ExtensionResourceRequest{
			Location: to.Ptr("Central US"),
			Plan: &armvisualstudio.ExtensionResourcePlan{
				Name:          to.Ptr("ExamplePlan"),
				Product:       to.Ptr("ExampleExtensionName"),
				PromotionCode: to.Ptr(""),
				Publisher:     to.Ptr("ExampleExtensionPublisher"),
				Version:       to.Ptr("1.0"),
			},
			Properties: map[string]*string{},
			Tags:       map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvisualstudio%2Farmvisualstudio%2Fv0.4.0/sdk/resourcemanager/visualstudio/armvisualstudio/README.md) on how to add the SDK to your project and authenticate.
