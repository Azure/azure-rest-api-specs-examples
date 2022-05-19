Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.6.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdatePython2Package.json
func ExamplePython2PackageClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewPython2PackageClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg",
		"myAutomationAccount33",
		"OmsCompositeResources",
		armautomation.PythonPackageCreateParameters{
			Properties: &armautomation.PythonPackageCreateProperties{
				ContentLink: &armautomation.ContentLink{
					ContentHash: &armautomation.ContentHash{
						Algorithm: to.Ptr("sha265"),
						Value:     to.Ptr("07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A"),
					},
					URI:     to.Ptr("https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip"),
					Version: to.Ptr("1.0.0.0"),
				},
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
