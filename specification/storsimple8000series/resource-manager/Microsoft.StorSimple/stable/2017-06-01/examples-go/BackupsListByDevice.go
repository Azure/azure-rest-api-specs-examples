package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsListByDevice.json
func ExampleBackupsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsClient().NewListByDevicePager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", &armstorsimple8000series.BackupsClientListByDeviceOptions{Filter: to.Ptr("createdTime%20ge%20'2017-06-22T18:30:00Z'%20and%20backupPolicyId%20eq%20'%2Fsubscriptions%2F4385cf00-2d3a-425a-832f-f4285b1c9dce%2FresourceGroups%2FResourceGroupForSDKTest%2Fproviders%2FMicrosoft.StorSimple%2Fmanagers%2FManagerForSDKTest1%2Fdevices%2FDevice05ForSDKTest%2FbackupPolicies%2FBkUpPolicy01ForSDKTest'")})
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
		// page.BackupList = armstorsimple8000series.BackupList{
		// 	Value: []*armstorsimple8000series.Backup{
		// 		{
		// 			Name: to.Ptr("880e1774-94a8-4f3e-85e6-a61e6b94a8b7"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/880e1774-94a8-4f3e-85e6-a61e6b94a8b7"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.BackupProperties{
		// 				BackupJobCreationType: to.Ptr(armstorsimple8000series.BackupJobCreationTypeAdhoc),
		// 				BackupPolicyID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest"),
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T22:01:19.708Z"); return t}()),
		// 				Elements: []*armstorsimple8000series.BackupElement{
		// 					{
		// 						ElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/880e1774-94a8-4f3e-85e6-a61e6b94a8b7/elements/87d398b9-63e8-4973-af85-12707c280ce8_0000000000000000"),
		// 						ElementName: to.Ptr("87d398b9-63e8-4973-af85-12707c280ce8_0000000000000000"),
		// 						ElementType: to.Ptr("managers/devices/backups/elements"),
		// 						SizeInBytes: to.Ptr[int64](10737418240),
		// 						VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 						VolumeName: to.Ptr("volume1"),
		// 						VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 					},
		// 					{
		// 						ElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/880e1774-94a8-4f3e-85e6-a61e6b94a8b7/elements/7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000"),
		// 						ElementName: to.Ptr("7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000"),
		// 						ElementType: to.Ptr("managers/devices/backups/elements"),
		// 						SizeInBytes: to.Ptr[int64](10737418240),
		// 						VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 						VolumeName: to.Ptr("Clonedvolume1"),
		// 						VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 				}},
		// 				SizeInBytes: to.Ptr[int64](21474836480),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("0826eda8-3d17-4cb9-b2af-d18ecf6ab819"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0826eda8-3d17-4cb9-b2af-d18ecf6ab819"),
		// 			Kind: to.Ptr("Series8000"),
		// 			Properties: &armstorsimple8000series.BackupProperties{
		// 				BackupJobCreationType: to.Ptr(armstorsimple8000series.BackupJobCreationTypeAdhoc),
		// 				BackupPolicyID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicyForSDKTest1032280949"),
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T21:22:01.372Z"); return t}()),
		// 				Elements: []*armstorsimple8000series.BackupElement{
		// 					{
		// 						ElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0826eda8-3d17-4cb9-b2af-d18ecf6ab819/elements/89f8af90-6420-45ec-9963-6b3e55d9a44c_0000000000000000"),
		// 						ElementName: to.Ptr("89f8af90-6420-45ec-9963-6b3e55d9a44c_0000000000000000"),
		// 						ElementType: to.Ptr("managers/devices/backups/elements"),
		// 						SizeInBytes: to.Ptr[int64](16106127360),
		// 						VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 						VolumeName: to.Ptr("volume2"),
		// 						VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 					},
		// 					{
		// 						ElementID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/0826eda8-3d17-4cb9-b2af-d18ecf6ab819/elements/95581ea2-b0aa-482a-be0c-8b7bd7658345_0000000000000000"),
		// 						ElementName: to.Ptr("95581ea2-b0aa-482a-be0c-8b7bd7658345_0000000000000000"),
		// 						ElementType: to.Ptr("managers/devices/backups/elements"),
		// 						SizeInBytes: to.Ptr[int64](10737418240),
		// 						VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
		// 						VolumeName: to.Ptr("volume1"),
		// 						VolumeType: to.Ptr(armstorsimple8000series.VolumeTypeTiered),
		// 				}},
		// 				SizeInBytes: to.Ptr[int64](26843545600),
		// 			},
		// 	}},
		// }
	}
}
