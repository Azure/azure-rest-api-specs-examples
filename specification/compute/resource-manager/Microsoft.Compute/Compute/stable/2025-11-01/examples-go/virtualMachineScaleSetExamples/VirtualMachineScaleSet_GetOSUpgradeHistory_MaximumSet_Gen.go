package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetOSUpgradeHistory_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewGetOSUpgradeHistoryPager_virtualMachineScaleSetGetOSUpgradeHistoryMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetsClient().NewGetOSUpgradeHistoryPager("rgcompute", "aaaaaa", nil)
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
		// page = armcompute.VirtualMachineScaleSetsClientGetOSUpgradeHistoryResponse{
		// 	VirtualMachineScaleSetListOSUpgradeHistory: armcompute.VirtualMachineScaleSetListOSUpgradeHistory{
		// 		Value: []*armcompute.UpgradeOperationHistoricalStatusInfo{
		// 			{
		// 				Properties: &armcompute.UpgradeOperationHistoricalStatusInfoProperties{
		// 					RunningStatus: &armcompute.UpgradeOperationHistoryStatus{
		// 						Code: to.Ptr(armcompute.UpgradeStateRollingForward),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:05:40.442Z"); return t}()),
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:05:40.443Z"); return t}()),
		// 					},
		// 					Progress: &armcompute.RollingUpgradeProgressInfo{
		// 						SuccessfulInstanceCount: to.Ptr[int32](6),
		// 						FailedInstanceCount: to.Ptr[int32](25),
		// 						InProgressInstanceCount: to.Ptr[int32](20),
		// 						PendingInstanceCount: to.Ptr[int32](27),
		// 					},
		// 					Error: &armcompute.APIError{
		// 						Details: []*armcompute.APIErrorBase{
		// 							{
		// 								Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Target: to.Ptr("aaaa"),
		// 								Message: to.Ptr("aa"),
		// 							},
		// 						},
		// 						Innererror: &armcompute.InnerError{
		// 							Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						Code: to.Ptr("aaaaaaa"),
		// 						Target: to.Ptr("aaaaaaa"),
		// 						Message: to.Ptr("aaaaaaaaa"),
		// 					},
		// 					StartedBy: to.Ptr(armcompute.UpgradeOperationInvokerUnknown),
		// 					TargetImageReference: &armcompute.ImageReference{
		// 						SKU: to.Ptr("2016-Datacenter"),
		// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 						Version: to.Ptr("latest"),
		// 						Offer: to.Ptr("WindowsServer"),
		// 						ExactVersion: to.Ptr("aaaaaaa"),
		// 						SharedGalleryImageID: to.Ptr("aaaaaa"),
		// 						ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					RollbackInfo: &armcompute.RollbackStatusInfo{
		// 						SuccessfullyRolledbackInstanceCount: to.Ptr[int32](12),
		// 						FailedRolledbackInstanceCount: to.Ptr[int32](2),
		// 						RollbackError: &armcompute.APIError{
		// 							Details: []*armcompute.APIErrorBase{
		// 								{
		// 									Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Target: to.Ptr("aaaa"),
		// 									Message: to.Ptr("aa"),
		// 								},
		// 							},
		// 							Innererror: &armcompute.InnerError{
		// 								Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 							Code: to.Ptr("aaaaaaa"),
		// 							Target: to.Ptr("aaaaaaa"),
		// 							Message: to.Ptr("aaaaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 				Location: to.Ptr("aaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaaaaaaa"),
		// 	},
		// }
	}
}
