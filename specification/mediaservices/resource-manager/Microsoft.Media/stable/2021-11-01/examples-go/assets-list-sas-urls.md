Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.4.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-sas-urls.json
func ExampleAssetsClient_ListContainerSas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewAssetsClient("<subscription-id>", cred, nil)
	res, err := client.ListContainerSas(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<asset-name>",
		armmediaservices.ListContainerSasInput{
			ExpiryTime:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T10:00:00.007Z"); return t }()),
			Permissions: armmediaservices.AssetContainerPermission("ReadWrite").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AssetsClientListContainerSasResult)
}
```
