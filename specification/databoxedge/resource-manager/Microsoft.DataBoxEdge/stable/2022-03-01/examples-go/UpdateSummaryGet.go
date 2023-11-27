package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UpdateSummaryGet.json
func ExampleDevicesClient_GetUpdateSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().GetUpdateSummary(ctx, "testedgedevice", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpdateSummary = armdataboxedge.UpdateSummary{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/updateSummary"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/updateSummary/default"),
	// 	Properties: &armdataboxedge.UpdateSummaryProperties{
	// 		DeviceLastScannedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-17T19:51:51.786Z"); return t}()),
	// 		DeviceVersionNumber: to.Ptr("2.1.1377.2170"),
	// 		FriendlyDeviceVersionName: to.Ptr("Azure Stack Edge 2010"),
	// 		LastCompletedScanJobDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-17T19:47:46.159Z"); return t}()),
	// 		LastDownloadJobStatus: to.Ptr(armdataboxedge.JobStatusInvalid),
	// 		LastInstallJobStatus: to.Ptr(armdataboxedge.JobStatusInvalid),
	// 		OngoingUpdateOperation: to.Ptr(armdataboxedge.UpdateOperationInstall),
	// 		RebootBehavior: to.Ptr(armdataboxedge.InstallRebootBehaviorNeverReboots),
	// 		TotalNumberOfUpdatesAvailable: to.Ptr[int32](1),
	// 		TotalNumberOfUpdatesPendingDownload: to.Ptr[int32](1),
	// 		TotalNumberOfUpdatesPendingInstall: to.Ptr[int32](1),
	// 		TotalUpdateSizeInBytes: to.Ptr[float64](4260898192),
	// 		UpdateTitles: []*string{
	// 			to.Ptr("Azure Stack Edge Update 2101 Package 1 of 2 for Pro-GPU, Pro R, Mini R")},
	// 		},
	// 	}
}
