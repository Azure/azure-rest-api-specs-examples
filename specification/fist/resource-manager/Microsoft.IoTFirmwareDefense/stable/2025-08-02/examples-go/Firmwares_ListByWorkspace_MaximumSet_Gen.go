package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-08-02/Firmwares_ListByWorkspace_MaximumSet_Gen.json
func ExampleFirmwaresClient_NewListByWorkspacePager_firmwaresListByWorkspaceMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirmwaresClient().NewListByWorkspacePager("rgiotfirmwaredefense", "exampleWorkspaceName", nil)
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
		// page = armiotfirmwaredefense.FirmwaresClientListByWorkspaceResponse{
		// 	FirmwareListResult: armiotfirmwaredefense.FirmwareListResult{
		// 		Value: []*armiotfirmwaredefense.Firmware{
		// 			{
		// 				Properties: &armiotfirmwaredefense.FirmwareProperties{
		// 					FileName: to.Ptr("FileNameThatWasUploaded.bin"),
		// 					Vendor: to.Ptr("ExampleVendorName"),
		// 					Model: to.Ptr("ExampleModelOfDevice"),
		// 					Version: to.Ptr("1.0.0"),
		// 					FileSize: to.Ptr[int64](30),
		// 					Status: to.Ptr(armiotfirmwaredefense.StatusPending),
		// 					StatusMessages: []*armiotfirmwaredefense.StatusMessage{
		// 						{
		// 							ErrorCode: to.Ptr[int64](5),
		// 							Message: to.Ptr("Firmware image contained some file systems that are not supported for extraction so results may be incomplete."),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 					Description: to.Ptr("User provided description of the firmware."),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/00000000-0000-0000-0000-000000000000"),
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares"),
		// 				SystemData: &armiotfirmwaredefense.SystemData{
		// 					CreatedBy: to.Ptr("UserName"),
		// 					CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserName"),
		// 					LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
