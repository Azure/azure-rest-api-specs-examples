package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksListByIscsiServer.json
func ExampleIscsiDisksClient_NewListByIscsiServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiDisksClient().NewListByIscsiServerPager("HSDK-0NZI14MDTF", "HSDK-0NZI14MDTF", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.ISCSIDiskList = armstorsimple1200series.ISCSIDiskList{
		// 	Value: []*armstorsimple1200series.ISCSIDisk{
		// 		{
		// 			Name: to.Ptr("Auto-TestIscsiDisk2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk2"),
		// 			Properties: &armstorsimple1200series.ISCSIDiskProperties{
		// 				Description: to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
		// 				AccessControlRecords: []*string{
		// 				},
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOnline),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Auto-TestIscsiDisk1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk1"),
		// 			Properties: &armstorsimple1200series.ISCSIDiskProperties{
		// 				Description: to.Ptr("Updated: Demo IscsiDisk for SDK Test Tiered"),
		// 				AccessControlRecords: []*string{
		// 				},
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOffline),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}
