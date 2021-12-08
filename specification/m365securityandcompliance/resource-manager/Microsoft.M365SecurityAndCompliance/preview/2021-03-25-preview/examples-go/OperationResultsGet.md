Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fm365securityandcompliance%2Farmm365securityandcompliance%2Fv0.1.0/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance/README.md) on how to add the SDK to your project and authenticate.

```go
package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/OperationResultsGet.json
func ExampleOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewOperationResultsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<location-name>",
		"<operation-result-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OperationResultsDescription.ID: %s\n", *res.ID)
}
```
