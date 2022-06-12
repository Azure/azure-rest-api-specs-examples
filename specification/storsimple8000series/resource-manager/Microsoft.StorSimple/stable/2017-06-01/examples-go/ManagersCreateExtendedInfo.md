```go
package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersCreateExtendedInfo.json
func ExampleManagersClient_CreateExtendedInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewManagersClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateExtendedInfo(ctx,
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest2",
		armstorsimple8000series.ManagerExtendedInfo{
			Properties: &armstorsimple8000series.ManagerExtendedInfoProperties{
				Algorithm:    to.Ptr("None"),
				IntegrityKey: to.Ptr("BIl+RHqO8PZ6DRvuXTTK7g=="),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple8000series%2Farmstorsimple8000series%2Fv1.0.0/sdk/resourcemanager/storsimple8000series/armstorsimple8000series/README.md) on how to add the SDK to your project and authenticate.
