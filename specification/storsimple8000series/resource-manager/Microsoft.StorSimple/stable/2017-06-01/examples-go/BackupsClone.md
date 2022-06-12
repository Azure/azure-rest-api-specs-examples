```go
package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsClone.json
func ExampleBackupsClient_BeginClone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginClone(ctx,
		"Device05ForSDKTest",
		"880e1774-94a8-4f3e-85e6-a61e6b94a8b7",
		"7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		armstorsimple8000series.CloneRequest{
			BackupElement: &armstorsimple8000series.BackupElement{
				ElementID:         to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/880e1774-94a8-4f3e-85e6-a61e6b94a8b7/elements/7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000"),
				ElementName:       to.Ptr("7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000"),
				ElementType:       to.Ptr("managers/devices/backups/elements"),
				SizeInBytes:       to.Ptr[int64](10737418240),
				VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
				VolumeName:        to.Ptr("Clonedvolume1"),
				VolumeType:        to.Ptr(armstorsimple8000series.VolumeTypeTiered),
			},
			TargetAccessControlRecordIDs: []*string{
				to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
			TargetDeviceID:   to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
			TargetVolumeName: to.Ptr("ClonedClonedvolume1"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple8000series%2Farmstorsimple8000series%2Fv1.0.0/sdk/resourcemanager/storsimple8000series/armstorsimple8000series/README.md) on how to add the SDK to your project and authenticate.
