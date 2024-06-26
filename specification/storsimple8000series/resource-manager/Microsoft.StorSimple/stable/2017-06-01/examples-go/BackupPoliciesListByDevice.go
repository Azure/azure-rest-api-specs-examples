package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupPoliciesListByDevice.json
func ExampleBackupPoliciesClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupPoliciesClient().NewListByDevicePager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
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
		// page.BackupPolicyList = armstorsimple8000series.BackupPolicyList{
		// 	Value: []*armstorsimple8000series.BackupPolicy{
		// 		{
		// 			Name: to.Ptr("BkUpPolicy01ForSDKTest"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.BackupPolicyProperties{
		// 				BackupPolicyCreationType: to.Ptr(armstorsimple8000series.BackupPolicyCreationTypeBySaaS),
		// 				LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T22:01:19.708Z"); return t}()),
		// 				NextBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T00:00:00.000Z"); return t}()),
		// 				ScheduledBackupStatus: to.Ptr(armstorsimple8000series.ScheduledBackupStatusEnabled),
		// 				SchedulesCount: to.Ptr[int64](2),
		// 				VolumeIDs: []*string{
		// 					to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/Clonedvolume1"),
		// 					to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("BkUpPolicyForSDKTest1032280949"),
		// 				Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies"),
		// 				ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicyForSDKTest1032280949"),
		// 				Kind: to.Ptr("Series8000"),
		// 				Properties: &armstorsimple8000series.BackupPolicyProperties{
		// 					BackupPolicyCreationType: to.Ptr(armstorsimple8000series.BackupPolicyCreationTypeBySaaS),
		// 					LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T21:22:01.372Z"); return t}()),
		// 					ScheduledBackupStatus: to.Ptr(armstorsimple8000series.ScheduledBackupStatusDisabled),
		// 					SchedulesCount: to.Ptr[int64](0),
		// 					VolumeIDs: []*string{
		// 						to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume2"),
		// 						to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1")},
		// 					},
		// 			}},
		// 		}
	}
}
