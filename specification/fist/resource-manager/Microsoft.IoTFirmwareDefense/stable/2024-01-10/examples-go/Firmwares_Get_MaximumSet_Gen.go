package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Firmwares_Get_MaximumSet_Gen.json
func ExampleFirmwaresClient_Get_firmwaresGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirmwaresClient().Get(ctx, "rgworkspaces-firmwares", "A7", "umrkdttp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Firmware = armiotfirmwaredefense.Firmware{
	// 	Name: to.Ptr("brmvnojpmxsgckdviynhxhftvcvbw"),
	// 	Type: to.Ptr("vlxhdbrphhnttmxvawwwewza"),
	// 	ID: to.Ptr("/subscriptions/blah/resourceGroups/blah/providers/blah/firmwares/blah"),
	// 	SystemData: &armiotfirmwaredefense.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T06:40:29.228Z"); return t}()),
	// 		CreatedBy: to.Ptr("aylilpicryidesowemwnl"),
	// 		CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T06:40:29.228Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bgtifu"),
	// 		LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 	},
	// 	Properties: &armiotfirmwaredefense.FirmwareProperties{
	// 		Description: to.Ptr("uz"),
	// 		FileName: to.Ptr("wresexxulcdsdd"),
	// 		FileSize: to.Ptr[int64](17),
	// 		Model: to.Ptr("f"),
	// 		ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
	// 		Status: to.Ptr(armiotfirmwaredefense.StatusPending),
	// 		StatusMessages: []*armiotfirmwaredefense.StatusMessage{
	// 			{
	// 				Message: to.Ptr("ulvhmhokezathzzauiitu"),
	// 		}},
	// 		Vendor: to.Ptr("vycmdhgtmepcptyoubztiuudpkcpd"),
	// 		Version: to.Ptr("s"),
	// 	},
	// }
}
