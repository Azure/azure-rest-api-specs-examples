package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksCreateOrUpdate.json
func ExampleIscsiDisksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewIscsiDisksClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"HSDK-0NZI14MDTF",
		"HSDK-0NZI14MDTF",
		"Auto-TestIscsiDisk1",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.ISCSIDisk{
			Name: to.Ptr("Auto-TestIscsiDisk1"),
			Properties: &armstorsimple1200series.ISCSIDiskProperties{
				Description:                to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
				AccessControlRecords:       []*string{},
				DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
				DiskStatus:                 to.Ptr(armstorsimple1200series.DiskStatusOnline),
				MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
