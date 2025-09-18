package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-08-02/Firmwares_Create_MaximumSet_Gen.json
func ExampleFirmwaresClient_Create_firmwaresCreateMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirmwaresClient().Create(ctx, "rgiotfirmwaredefense", "exampleWorkspaceName", "00000000-0000-0000-0000-000000000000", armiotfirmwaredefense.Firmware{
		Properties: &armiotfirmwaredefense.FirmwareProperties{
			FileName: to.Ptr("NameOfFirmwareFile.bin"),
			Vendor:   to.Ptr("ExampleVendorName"),
			Model:    to.Ptr("ExampleModelOfDevice"),
			Version:  to.Ptr("1.0.0"),
			FileSize: to.Ptr[int64](30),
			Status:   to.Ptr(armiotfirmwaredefense.StatusPending),
			StatusMessages: []*armiotfirmwaredefense.StatusMessage{
				{
					ErrorCode: to.Ptr[int64](200),
					Message:   to.Ptr("What would this message say?"),
				},
			},
			Description: to.Ptr("User provided description of the firmware."),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotfirmwaredefense.FirmwaresClientCreateResponse{
	// 	Firmware: &armiotfirmwaredefense.Firmware{
	// 		Properties: &armiotfirmwaredefense.FirmwareProperties{
	// 			FileName: to.Ptr("FileNameThatWasUploaded.bin"),
	// 			Vendor: to.Ptr("ExampleVendorName"),
	// 			Model: to.Ptr("ExampleModelOfDevice"),
	// 			Version: to.Ptr("1.0.0"),
	// 			FileSize: to.Ptr[int64](30),
	// 			Status: to.Ptr(armiotfirmwaredefense.StatusPending),
	// 			StatusMessages: []*armiotfirmwaredefense.StatusMessage{
	// 				{
	// 					ErrorCode: to.Ptr[int64](200),
	// 					Message: to.Ptr("Still dont know what this message would say."),
	// 				},
	// 			},
	// 			Description: to.Ptr("User provided description of the firmware."),
	// 			ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ResourceGroupName/providers/Microsoft.IoTFirmwareDefense/workspaces/WorkspaceName/firmwares/00000000-0000-0000-0000-000000000000"),
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares"),
	// 		SystemData: &armiotfirmwaredefense.SystemData{
	// 			CreatedBy: to.Ptr("UserName"),
	// 			CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserName"),
	// 			LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
	// 		},
	// 	},
	// }
}
