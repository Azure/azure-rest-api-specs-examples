package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesCreateOrUpdate.json
func ExampleFileSharesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSharesClient().BeginCreateOrUpdate(ctx, "HSDK-4XY4FI2IVG", "HSDK-4XY4FI2IVG", "Auto-TestFileShare1", "ResourceGroupForSDKTest", "hAzureSDKOperations", armstorsimple1200series.FileShare{
		Name: to.Ptr("Auto-TestFileShare1"),
		Properties: &armstorsimple1200series.FileShareProperties{
			Description:                to.Ptr("Demo FileShare for SDK Test Tiered"),
			AdminUser:                  to.Ptr("fareast\\idcdlslb"),
			DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
			MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
			ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
			ShareStatus:                to.Ptr(armstorsimple1200series.ShareStatusOnline),
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
	// res.FileShare = armstorsimple1200series.FileShare{
	// 	Name: to.Ptr("Auto-TestFileShare1"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/fileServers/shares"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/fileServers/HSDK-4XY4FI2IVG/shares/Auto-TestFileShare1"),
	// 	Properties: &armstorsimple1200series.FileShareProperties{
	// 		Description: to.Ptr("Demo FileShare for SDK Test Tiered"),
	// 		AdminUser: to.Ptr("fareast\\idcdlslb"),
	// 		DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
	// 		LocalUsedCapacityInBytes: to.Ptr[int64](0),
	// 		MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
	// 		ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
	// 		ShareStatus: to.Ptr(armstorsimple1200series.ShareStatusOnline),
	// 		UsedCapacityInBytes: to.Ptr[int64](0),
	// 	},
	// }
}
