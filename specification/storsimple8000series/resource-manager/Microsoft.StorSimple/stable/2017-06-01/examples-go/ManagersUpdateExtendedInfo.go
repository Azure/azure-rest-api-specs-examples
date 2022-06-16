package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersUpdateExtendedInfo.json
func ExampleManagersClient_UpdateExtendedInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewManagersClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateExtendedInfo(ctx,
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest2",
		"674ab9df-af56-4c5c-a5ca-2bddcf1b781a",
		armstorsimple8000series.ManagerExtendedInfo{
			Etag: to.Ptr("674ab9df-af56-4c5c-a5ca-2bddcf1b781a"),
			Properties: &armstorsimple8000series.ManagerExtendedInfoProperties{
				Algorithm:    to.Ptr("SHA256"),
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
