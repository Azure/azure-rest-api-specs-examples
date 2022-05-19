Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsesu%2Farmwindowsesu%2Fv0.5.0/sdk/resourcemanager/windowsesu/armwindowsesu/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsesu_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/CreateMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armwindowsesu.NewMultipleActivationKeysClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testgr1",
		"server08-key-2019",
		armwindowsesu.MultipleActivationKey{
			Location: to.Ptr("East US"),
			Properties: &armwindowsesu.MultipleActivationKeyProperties{
				AgreementNumber:       to.Ptr("1a2b45ag"),
				InstalledServerNumber: to.Ptr[int32](100),
				IsEligible:            to.Ptr(true),
				OSType:                to.Ptr(armwindowsesu.OsTypeWindowsServer2008),
				SupportType:           to.Ptr(armwindowsesu.SupportTypeSupplementalServicing),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
