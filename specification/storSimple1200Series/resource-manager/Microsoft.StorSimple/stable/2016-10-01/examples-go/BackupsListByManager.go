package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/BackupsListByManager.json
func ExampleBackupsClient_NewListByManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsClient().NewListByManagerPager("ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.BackupsClientListByManagerOptions{Filter: to.Ptr("createdTime%20ge%20'2018-08-10T17:30:03Z'%20and%20createdTime%20le%20'2018-08-14T17:30:03Z'%20and%20initiatedBy%20eq%20'Manual'")})
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
		// page.BackupList = armstorsimple1200series.BackupList{
		// 	Value: []*armstorsimple1200series.Backup{
		// 		{
		// 			Name: to.Ptr("315c3461-96ad-41bf-af0b-3a8be5113203"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/315c3461-96ad-41bf-af0b-3a8be5113203"),
		// 			Properties: &armstorsimple1200series.BackupProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T14:34:52.414Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI"),
		// 				Elements: []*armstorsimple1200series.BackupElement{
		// 					{
		// 						Name: to.Ptr("f3fa955c-e20e-46fd-a71e-c0f0491db14b"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/315c3461-96ad-41bf-af0b-3a8be5113203/elements/f3fa955c-e20e-46fd-a71e-c0f0491db14b"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestIscsiDisk2"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("9ba7098e-0fc8-43de-b39b-a4228dd3bbde"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/315c3461-96ad-41bf-af0b-3a8be5113203/elements/9ba7098e-0fc8-43de-b39b-a4228dd3bbde"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestIscsiDisk1"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 				}},
		// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				InitiatedBy: to.Ptr(armstorsimple1200series.InitiatedByManual),
		// 				SizeInBytes: to.Ptr[int64](1073741824000),
		// 				TargetID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/iscsiServers/HSDK-UGU4PITWNI"),
		// 				TargetType: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("cce3a13c-b37e-4db1-bb78-444cc85be33d"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/cce3a13c-b37e-4db1-bb78-444cc85be33d"),
		// 			Properties: &armstorsimple1200series.BackupProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-12T14:02:13.067Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI"),
		// 				Elements: []*armstorsimple1200series.BackupElement{
		// 					{
		// 						Name: to.Ptr("bebfaa5b-b02f-4f14-bf3f-83a0abdff090"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/cce3a13c-b37e-4db1-bb78-444cc85be33d/elements/bebfaa5b-b02f-4f14-bf3f-83a0abdff090"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("ClonedTieredIscsiDiskForSDKTest"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("8e065d31-0571-43cc-a134-4855603bc222"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/cce3a13c-b37e-4db1-bb78-444cc85be33d/elements/8e065d31-0571-43cc-a134-4855603bc222"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestIscsiDisk2"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("e4a4b2c8-0252-488c-8909-f7bc67474b62"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/backups/cce3a13c-b37e-4db1-bb78-444cc85be33d/elements/e4a4b2c8-0252-488c-8909-f7bc67474b62"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestIscsiDisk1"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 				}},
		// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				InitiatedBy: to.Ptr(armstorsimple1200series.InitiatedByManual),
		// 				SizeInBytes: to.Ptr[int64](1610612736000),
		// 				TargetID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-UGU4PITWNI/iscsiServers/HSDK-UGU4PITWNI"),
		// 				TargetType: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("899222aa-1340-4090-a8b1-7436e2b859d3"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/899222aa-1340-4090-a8b1-7436e2b859d3"),
		// 			Properties: &armstorsimple1200series.BackupProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-12T14:34:24.398Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0"),
		// 				Elements: []*armstorsimple1200series.BackupElement{
		// 					{
		// 						Name: to.Ptr("ac877a33-cce2-4107-a3f5-b00e0af43bd8"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/899222aa-1340-4090-a8b1-7436e2b859d3/elements/ac877a33-cce2-4107-a3f5-b00e0af43bd8"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyLocal),
		// 							EndpointName: to.Ptr("Auto-TestFileShare2"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("6c5afd0b-5b1f-47c5-a37e-eab0812cfdf3"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/899222aa-1340-4090-a8b1-7436e2b859d3/elements/6c5afd0b-5b1f-47c5-a37e-eab0812cfdf3"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("ClonedTieredFileShareForSDKTest"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("6565f0f0-97d2-427f-8e44-fc43c4558234"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/899222aa-1340-4090-a8b1-7436e2b859d3/elements/6565f0f0-97d2-427f-8e44-fc43c4558234"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestFileShare1"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 				}},
		// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				InitiatedBy: to.Ptr(armstorsimple1200series.InitiatedByManual),
		// 				SizeInBytes: to.Ptr[int64](1610612736000),
		// 				TargetID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/fileServers/HSDK-DMNJB2PET0"),
		// 				TargetType: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a4ec37c8-7e5e-4483-88ec-9b37fdd686ff"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/a4ec37c8-7e5e-4483-88ec-9b37fdd686ff"),
		// 			Properties: &armstorsimple1200series.BackupProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-12T13:40:04.590Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0"),
		// 				Elements: []*armstorsimple1200series.BackupElement{
		// 					{
		// 						Name: to.Ptr("14e4f48f-e3fe-40cd-a3f4-5b226da83278"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/a4ec37c8-7e5e-4483-88ec-9b37fdd686ff/elements/14e4f48f-e3fe-40cd-a3f4-5b226da83278"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyLocal),
		// 							EndpointName: to.Ptr("Auto-TestFileShare2"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("bb6ec523-f18c-4123-8e44-ce1e83db1ab1"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/a4ec37c8-7e5e-4483-88ec-9b37fdd686ff/elements/bb6ec523-f18c-4123-8e44-ce1e83db1ab1"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("ClonedTieredFileShareForSDKTest"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("dea75831-c6bf-4ae9-9f6f-3b6eae680db5"),
		// 						Type: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements"),
		// 						ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/backups/a4ec37c8-7e5e-4483-88ec-9b37fdd686ff/elements/dea75831-c6bf-4ae9-9f6f-3b6eae680db5"),
		// 						Properties: &armstorsimple1200series.BackupElementProperties{
		// 							DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 							EndpointName: to.Ptr("Auto-TestFileShare1"),
		// 							SizeInBytes: to.Ptr[int64](536870912000),
		// 						},
		// 				}},
		// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				InitiatedBy: to.Ptr(armstorsimple1200series.InitiatedByManual),
		// 				SizeInBytes: to.Ptr[int64](1610612736000),
		// 				TargetID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-DMNJB2PET0/fileServers/HSDK-DMNJB2PET0"),
		// 				TargetType: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers"),
		// 			},
		// 	}},
		// }
	}
}
