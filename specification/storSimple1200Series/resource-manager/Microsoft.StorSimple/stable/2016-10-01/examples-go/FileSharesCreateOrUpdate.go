package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesCreateOrUpdate.json
func ExampleFileSharesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewFileSharesClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"HSDK-4XY4FI2IVG",
		"HSDK-4XY4FI2IVG",
		"Auto-TestFileShare1",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.FileShare{
			Name: to.Ptr("Auto-TestFileShare1"),
			Properties: &armstorsimple1200series.FileShareProperties{
				Description:                to.Ptr("Demo FileShare for SDK Test Tiered"),
				AdminUser:                  to.Ptr("fareast\\idcdlslb"),
				DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
				MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
				ShareStatus:                to.Ptr(armstorsimple1200series.ShareStatusOnline),
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
