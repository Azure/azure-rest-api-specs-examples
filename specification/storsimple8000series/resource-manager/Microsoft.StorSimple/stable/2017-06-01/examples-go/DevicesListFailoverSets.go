package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesListFailoverSets.json
func ExampleDevicesClient_NewListFailoverSetsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListFailoverSetsPager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FailoverSetsList = armstorsimple8000series.FailoverSetsList{
		// 	Value: []*armstorsimple8000series.FailoverSet{
		// 		{
		// 			EligibilityResult: &armstorsimple8000series.FailoverSetEligibilityResult{
		// 				IsEligibleForFailover: to.Ptr(true),
		// 			},
		// 			VolumeContainers: []*armstorsimple8000series.VolumeContainerFailoverMetadata{
		// 				{
		// 					VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcforsdktest"),
		// 					Volumes: []*armstorsimple8000series.VolumeFailoverMetadata{
		// 						{
		// 							BackupCreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T06:30:03.794Z"); return t}()),
		// 							BackupElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/964d5a65-5294-4707-9c71-becb8850ea12/elements/2553386f-f39e-4223-b9fa-319adc5630fe_0000000000000000"),
		// 							BackupID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/964d5a65-5294-4707-9c71-becb8850ea12"),
		// 							BackupPolicyID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/jembkpolicy"),
		// 							SizeInBytes: to.Ptr[int64](5368709120),
		// 							VolumeID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcforsdktest/volumes/jemVol1"),
		// 							VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			EligibilityResult: &armstorsimple8000series.FailoverSetEligibilityResult{
		// 				IsEligibleForFailover: to.Ptr(true),
		// 			},
		// 			VolumeContainers: []*armstorsimple8000series.VolumeContainerFailoverMetadata{
		// 				{
		// 					VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcForOdataFilterTest"),
		// 					Volumes: []*armstorsimple8000series.VolumeFailoverMetadata{
		// 						{
		// 							BackupCreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-10T10:29:37.215Z"); return t}()),
		// 							BackupElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0874889e-f2a3-42d2-a643-f378fc199688/elements/5baae396-46f1-46a0-bc25-d8e2fcada7fa_0000000000000000"),
		// 							BackupID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0874889e-f2a3-42d2-a643-f378fc199688"),
		// 							BackupPolicyID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicyForSDKTest429326237"),
		// 							SizeInBytes: to.Ptr[int64](10737418240),
		// 							VolumeID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcForOdataFilterTest/volumes/vol01"),
		// 							VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 						},
		// 						{
		// 							BackupCreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-10T10:29:37.215Z"); return t}()),
		// 							BackupElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0874889e-f2a3-42d2-a643-f378fc199688/elements/5c7eb677-d688-4afb-a352-ca62454d5921_0000000000000000"),
		// 							BackupID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0874889e-f2a3-42d2-a643-f378fc199688"),
		// 							BackupPolicyID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicyForSDKTest429326237"),
		// 							SizeInBytes: to.Ptr[int64](10737418240),
		// 							VolumeID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcForOdataFilterTest/volumes/vol%2540123"),
		// 							VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 					}},
		// 			}},
		// 	}},
		// }
	}
}
