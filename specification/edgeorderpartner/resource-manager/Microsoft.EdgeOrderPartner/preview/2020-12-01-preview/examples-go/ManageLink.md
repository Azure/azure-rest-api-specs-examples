Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorderpartner%2Farmedgeorderpartner%2Fv0.2.1/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner/README.md) on how to add the SDK to your project and authenticate.

```go
package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// x-ms-original-file: specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageLink.json
func ExampleAPISClient_ManageLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorderpartner.NewAPISClient("<subscription-id>", cred, nil)
	_, err = client.ManageLink(ctx,
		"<family-identifier>",
		"<location>",
		"<serial-number>",
		armedgeorderpartner.ManageLinkRequest{
			ManagementResourceArmID: to.StringPtr("<management-resource-arm-id>"),
			Operation:               armedgeorderpartner.ManageLinkOperationLink.ToPtr(),
			TenantID:                to.StringPtr("<tenant-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
