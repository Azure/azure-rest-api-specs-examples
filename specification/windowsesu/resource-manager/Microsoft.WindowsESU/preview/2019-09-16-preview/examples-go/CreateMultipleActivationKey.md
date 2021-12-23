Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsesu%2Farmwindowsesu%2Fv0.1.0/sdk/resourcemanager/windowsesu/armwindowsesu/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsesu_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"
)

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/CreateMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<multiple-activation-key-name>",
		armwindowsesu.MultipleActivationKey{
			TrackedResource: armwindowsesu.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armwindowsesu.MultipleActivationKeyProperties{
				AgreementNumber:       to.StringPtr("<agreement-number>"),
				InstalledServerNumber: to.Int32Ptr(100),
				IsEligible:            to.BoolPtr(true),
				OSType:                armwindowsesu.OsTypeWindowsServer2008.ToPtr(),
				SupportType:           armwindowsesu.SupportTypeSupplementalServicing.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MultipleActivationKey.ID: %s\n", *res.ID)
}
```
