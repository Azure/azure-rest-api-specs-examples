```go
package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersUpdateExtendedInfo.json
func ExampleManagersClient_UpdateExtendedInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewManagersClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateExtendedInfo(ctx,
		"ResourceGroupForSDKTest",
		"hManagerForSDKTest",
		"6531d5d7-3ced-4f78-83b6-804368f2ca0c",
		armstorsimple1200series.ManagerExtendedInfo{
			Name: to.Ptr("vaultExtendedInfo"),
			Type: to.Ptr("Microsoft.StorSimple/Managers/extendedInformation"),
			ID:   to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/hManagerForSDKTestextendedInformation/vaultExtendedInfo"),
			Etag: to.Ptr("6531d5d7-3ced-4f78-83b6-804368f2ca0c"),
			Properties: &armstorsimple1200series.ManagerExtendedInfoProperties{
				Algorithm:    to.Ptr("SHA256"),
				IntegrityKey: to.Ptr("e6501980-7efe-4602-bb0e-3cb9a08a6003"),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple1200series%2Farmstorsimple1200series%2Fv1.0.0/sdk/resourcemanager/storsimple1200series/armstorsimple1200series/README.md) on how to add the SDK to your project and authenticate.
