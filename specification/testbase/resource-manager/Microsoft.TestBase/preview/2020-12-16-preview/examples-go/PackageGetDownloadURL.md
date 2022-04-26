Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.4.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/PackageGetDownloadURL.json
func ExamplePackagesClient_GetDownloadURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armtestbase.NewPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetDownloadURL(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
