Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.1.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/CheckPackageNameAvailability.json
func ExampleTestBaseAccountsClient_CheckPackageNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestBaseAccountsClient("<subscription-id>", cred, nil)
	_, err = client.CheckPackageNameAvailability(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		armtestbase.PackageCheckNameAvailabilityParameters{
			Name:            to.StringPtr("<name>"),
			Type:            to.StringPtr("<type>"),
			ApplicationName: to.StringPtr("<application-name>"),
			Version:         to.StringPtr("<version>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
