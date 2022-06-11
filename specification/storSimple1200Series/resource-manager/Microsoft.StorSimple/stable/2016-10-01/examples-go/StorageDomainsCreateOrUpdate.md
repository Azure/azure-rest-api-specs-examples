```go
package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/StorageDomainsCreateOrUpdate.json
func ExampleStorageDomainsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewStorageDomainsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"sd-fs-HSDK-4XY4FI2IVG",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.StorageDomain{
			Name: to.Ptr("sd-fs-HSDK-4XY4FI2IVG"),
			Properties: &armstorsimple1200series.StorageDomainProperties{
				EncryptionStatus: to.Ptr(armstorsimple1200series.EncryptionStatusDisabled),
				StorageAccountCredentialIDs: []*string{
					to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageAccountCredentials/sacforsdktest")},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple1200series%2Farmstorsimple1200series%2Fv1.0.0/sdk/resourcemanager/storsimple1200series/armstorsimple1200series/README.md) on how to add the SDK to your project and authenticate.
