package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Summaries_ListByFirmware_MaximumSet_Gen.json
func ExampleSummariesClient_NewListByFirmwarePager_summariesListByFirmwareMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSummariesClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
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
		// page.SummaryListResult = armiotfirmwaredefense.SummaryListResult{
		// 	Value: []*armiotfirmwaredefense.SummaryResource{
		// 		{
		// 			Name: to.Ptr("CVE"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/summaries"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/CVE"),
		// 			Properties: &armiotfirmwaredefense.CveSummary{
		// 				SummaryType: to.Ptr(armiotfirmwaredefense.SummaryTypeCVE),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("BinaryHardening"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/summaries"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/BinaryHardening"),
		// 			Properties: &armiotfirmwaredefense.BinaryHardeningSummaryResource{
		// 				SummaryType: to.Ptr(armiotfirmwaredefense.SummaryTypeBinaryHardening),
		// 			},
		// 	}},
		// }
	}
}
