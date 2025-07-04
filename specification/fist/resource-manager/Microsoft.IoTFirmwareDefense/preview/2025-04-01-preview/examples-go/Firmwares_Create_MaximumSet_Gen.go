package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Firmwares_Create_MaximumSet_Gen.json
func ExampleFirmwaresClient_Create_firmwaresCreateMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("5C707B5F-6130-4F71-819E-953A28942E88", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirmwaresClient().Create(ctx, "rgiotfirmwaredefense", "exampleWorkspaceName", "00000000-0000-0000-0000-000000000000", armiotfirmwaredefense.Firmware{
		Properties: &armiotfirmwaredefense.FirmwareProperties{
			FileName: to.Ptr("dmnqhyxssutvnewntlb"),
			Vendor:   to.Ptr("hymojocxpxqhtblioaavylnzyg"),
			Model:    to.Ptr("wmyfbyjsggbvxcuin"),
			Version:  to.Ptr("nhtxzslgcbtptu"),
			FileSize: to.Ptr[int64](30),
			Status:   to.Ptr(armiotfirmwaredefense.StatusPending),
			StatusMessages: []*armiotfirmwaredefense.StatusMessage{
				{
					ErrorCode: to.Ptr[int64](20),
					Message:   to.Ptr("edtylkjvj"),
				},
			},
			Description: to.Ptr("sqt"),
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
	// 			FileName: to.Ptr("dmnqhyxssutvnewntlb"),
	// 			Vendor: to.Ptr("hymojocxpxqhtblioaavylnzyg"),
	// 			Model: to.Ptr("wmyfbyjsggbvxcuin"),
	// 			Version: to.Ptr("nhtxzslgcbtptu"),
	// 			FileSize: to.Ptr[int64](30),
	// 			Status: to.Ptr(armiotfirmwaredefense.StatusPending),
	// 			StatusMessages: []*armiotfirmwaredefense.StatusMessage{
	// 				{
	// 					ErrorCode: to.Ptr[int64](20),
	// 					Message: to.Ptr("edtylkjvj"),
	// 				},
	// 			},
	// 			Description: to.Ptr("sqt"),
	// 			ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/firmware"),
	// 		Name: to.Ptr("qobb"),
	// 		Type: to.Ptr("xf"),
	// 		SystemData: &armiotfirmwaredefense.SystemData{
	// 			CreatedBy: to.Ptr("nqisshvdzqcxzbujvacin"),
	// 			CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("of"),
	// 			LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
	// 		},
	// 	},
	// }
}
