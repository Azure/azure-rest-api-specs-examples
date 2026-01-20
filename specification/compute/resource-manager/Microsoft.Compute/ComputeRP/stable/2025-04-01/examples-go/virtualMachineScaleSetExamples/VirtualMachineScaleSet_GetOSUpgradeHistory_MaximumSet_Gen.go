package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetOSUpgradeHistory_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewGetOSUpgradeHistoryPager_virtualMachineScaleSetGetOsUpgradeHistoryMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.VirtualMachineScaleSetListOSUpgradeHistory = armcompute.VirtualMachineScaleSetListOSUpgradeHistory{
		// 	Value: []*armcompute.UpgradeOperationHistoricalStatusInfo{
		// 		{
		// 			Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 			Location: to.Ptr("aaaaaaaaaaaaa"),
		// 			Properties: &armcompute.UpgradeOperationHistoricalStatusInfoProperties{
		// 				Error: &armcompute.APIError{
		// 					Code: to.Ptr("aaaaaaa"),
		// 					Innererror: &armcompute.InnerError{
		// 						Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					Message: to.Ptr("aaaaaaaaa"),
		// 					Target: to.Ptr("aaaaaaa"),
		// 					Details: []*armcompute.APIErrorBase{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Message: to.Ptr("aa"),
		// 							Target: to.Ptr("aaaa"),
		// 					}},
		// 				},
		// 				Progress: &armcompute.RollingUpgradeProgressInfo{
		// 					FailedInstanceCount: to.Ptr[int32](25),
		// 					InProgressInstanceCount: to.Ptr[int32](20),
		// 					PendingInstanceCount: to.Ptr[int32](27),
		// 					SuccessfulInstanceCount: to.Ptr[int32](6),
		// 				},
		// 				RollbackInfo: &armcompute.RollbackStatusInfo{
		// 					FailedRolledbackInstanceCount: to.Ptr[int32](2),
		// 					RollbackError: &armcompute.APIError{
		// 						Code: to.Ptr("aaaaaaa"),
		// 						Innererror: &armcompute.InnerError{
		// 							Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						Message: to.Ptr("aaaaaaaaa"),
		// 						Target: to.Ptr("aaaaaaa"),
		// 						Details: []*armcompute.APIErrorBase{
		// 							{
		// 								Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Message: to.Ptr("aa"),
		// 								Target: to.Ptr("aaaa"),
		// 						}},
		// 					},
		// 					SuccessfullyRolledbackInstanceCount: to.Ptr[int32](12),
		// 				},
		// 				RunningStatus: &armcompute.UpgradeOperationHistoryStatus{
		// 					Code: to.Ptr(armcompute.UpgradeStateRollingForward),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:05:40.443Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:05:40.442Z"); return t}()),
		// 				},
		// 				StartedBy: to.Ptr(armcompute.UpgradeOperationInvokerUnknown),
		// 				TargetImageReference: &armcompute.ImageReference{
		// 					ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 					ExactVersion: to.Ptr("aaaaaaa"),
		// 					Offer: to.Ptr("WindowsServer"),
		// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 					SharedGalleryImageID: to.Ptr("aaaaaa"),
		// 					SKU: to.Ptr("2016-Datacenter"),
		// 					Version: to.Ptr("latest"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
