package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/SbomComponents_ListByFirmware_MaximumSet_Gen.json
func ExampleSbomComponentsClient_NewListByFirmwarePager_sbomComponentsListByFirmwareMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSbomComponentsClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
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
		// page.SbomComponentListResult = armiotfirmwaredefense.SbomComponentListResult{
		// 	Value: []*armiotfirmwaredefense.SbomComponentResource{
		// 		{
		// 			Name: to.Ptr("cd31ca40-6772-44f8-9180-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/sbomComponents"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/sbomComponents/cd31ca40-6772-44f8-9180-000000000000"),
		// 			Properties: &armiotfirmwaredefense.SbomComponent{
		// 				ComponentID: to.Ptr("busybox:1.30.1"),
		// 				ComponentName: to.Ptr("busybox"),
		// 				FilePaths: []*string{
		// 					to.Ptr("/path/to/bin/busybox")},
		// 					License: to.Ptr("GPL-2.0-only"),
		// 					Version: to.Ptr("1.30.1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("5924e71d-5f38-444e-8b08-000000000000"),
		// 				Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/sbomComponents"),
		// 				ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/sbomComponents/5924e71d-5f38-444e-8b08-000000000000"),
		// 				Properties: &armiotfirmwaredefense.SbomComponent{
		// 					ComponentID: to.Ptr("openvpn:2.4.6"),
		// 					ComponentName: to.Ptr("openvpn"),
		// 					FilePaths: []*string{
		// 						to.Ptr("/path/to/squashfs-root/bin/openvpn")},
		// 						License: to.Ptr("GPL-2.0"),
		// 						Version: to.Ptr("2.4.6"),
		// 					},
		// 			}},
		// 		}
	}
}
