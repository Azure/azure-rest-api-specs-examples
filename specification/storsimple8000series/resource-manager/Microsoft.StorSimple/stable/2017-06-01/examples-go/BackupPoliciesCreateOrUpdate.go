package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupPoliciesCreateOrUpdate.json
func ExampleBackupPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupPoliciesClient().BeginCreateOrUpdate(ctx, "Device05ForSDKTest", "BkUpPolicy01ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.BackupPolicy{
		Kind: to.Ptr("Series8000"),
		Properties: &armstorsimple8000series.BackupPolicyProperties{
			VolumeIDs: []*string{
				to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/Clonedvolume1"),
				to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupPolicy = armstorsimple8000series.BackupPolicy{
	// 	Name: to.Ptr("BkUpPolicy01ForSDKTest"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.BackupPolicyProperties{
	// 		BackupPolicyCreationType: to.Ptr(armstorsimple8000series.BackupPolicyCreationTypeBySaaS),
	// 		ScheduledBackupStatus: to.Ptr(armstorsimple8000series.ScheduledBackupStatusDisabled),
	// 		SchedulesCount: to.Ptr[int64](0),
	// 		VolumeIDs: []*string{
	// 			to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/Clonedvolume1"),
	// 			to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1")},
	// 		},
	// 	}
}
