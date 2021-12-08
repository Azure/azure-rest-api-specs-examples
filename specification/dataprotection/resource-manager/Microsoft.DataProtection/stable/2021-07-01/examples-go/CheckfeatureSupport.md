Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataprotection%2Farmdataprotection%2Fv0.1.0/sdk/resourcemanager/dataprotection/armdataprotection/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/CheckfeatureSupport.json
func ExampleDataProtectionClient_CheckFeatureSupport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewDataProtectionClient("<subscription-id>", cred, nil)
	_, err = client.CheckFeatureSupport(ctx,
		"<location>",
		&armdataprotection.FeatureValidationRequest{
			FeatureValidationRequestBase: armdataprotection.FeatureValidationRequestBase{
				ObjectType: to.StringPtr("<object-type>"),
			},
			FeatureType: armdataprotection.FeatureTypeDataSourceType.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
