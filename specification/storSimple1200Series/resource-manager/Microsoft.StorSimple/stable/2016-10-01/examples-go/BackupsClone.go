package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/BackupsClone.json
func ExampleBackupsClient_BeginClone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewBackupsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginClone(ctx,
		"HSDK-4XY4FI2IVG",
		"58d33025-e837-41cc-b15f-7c85ced64aab",
		"2304968f-91af-4f59-8b79-31e321eee40e",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.CloneRequest{
			Properties: &armstorsimple1200series.CloneRequestProperties{
				NewEndpointName: to.Ptr("ClonedTieredFileShareForSDKTest"),
				Share: &armstorsimple1200series.FileShare{
					Name: to.Ptr("TieredFileShareForSDKTest"),
					Properties: &armstorsimple1200series.FileShareProperties{
						Description:                to.Ptr("Restore file Share"),
						AdminUser:                  to.Ptr("HSDK-4XY4FI2IVG\\StorSimpleAdmin"),
						DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
						MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
						ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
						ShareStatus:                to.Ptr(armstorsimple1200series.ShareStatusOnline),
					},
				},
				TargetAccessPointID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG"),
				TargetDeviceID:      to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg"),
			},
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
