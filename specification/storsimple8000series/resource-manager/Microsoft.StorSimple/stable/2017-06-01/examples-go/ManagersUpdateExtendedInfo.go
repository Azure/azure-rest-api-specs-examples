package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersUpdateExtendedInfo.json
func ExampleManagersClient_UpdateExtendedInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().UpdateExtendedInfo(ctx, "ResourceGroupForSDKTest", "ManagerForSDKTest2", "674ab9df-af56-4c5c-a5ca-2bddcf1b781a", armstorsimple8000series.ManagerExtendedInfo{
		Etag: to.Ptr("674ab9df-af56-4c5c-a5ca-2bddcf1b781a"),
		Properties: &armstorsimple8000series.ManagerExtendedInfoProperties{
			Algorithm:    to.Ptr("SHA256"),
			IntegrityKey: to.Ptr("BIl+RHqO8PZ6DRvuXTTK7g=="),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerExtendedInfo = armstorsimple8000series.ManagerExtendedInfo{
	// 	Name: to.Ptr("vaultExtendedInfo"),
	// 	Type: to.Ptr("Microsoft.StorSimple/Managers/extendedInformation"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/Managers/ManagerForSDKTest2extendedInformation/vaultExtendedInfo"),
	// 	Etag: to.Ptr("7d062cd9-c325-422b-8bff-0e702b8659ed"),
	// 	Properties: &armstorsimple8000series.ManagerExtendedInfoProperties{
	// 		Algorithm: to.Ptr("SHA256"),
	// 		IntegrityKey: to.Ptr("BIl+RHqO8PZ6DRvuXTTK7g=="),
	// 	},
	// }
}
