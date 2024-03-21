package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Summaries_Get_MaximumSet_Gen.json
func ExampleSummariesClient_Get_summariesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSummariesClient().Get(ctx, "FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", armiotfirmwaredefense.SummaryNameFirmware, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SummaryResource = armiotfirmwaredefense.SummaryResource{
	// 	Name: to.Ptr("firmware"),
	// 	Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/summaries"),
	// 	ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/firmware"),
	// 	Properties: &armiotfirmwaredefense.FirmwareSummary{
	// 		SummaryType: to.Ptr(armiotfirmwaredefense.SummaryTypeFirmware),
	// 		AnalysisTimeSeconds: to.Ptr[int64](172),
	// 		BinaryCount: to.Ptr[int64](579),
	// 		ComponentCount: to.Ptr[int64](35),
	// 		ExtractedFileCount: to.Ptr[int64](4171),
	// 		ExtractedSize: to.Ptr[int64](302752916),
	// 		FileSize: to.Ptr[int64](55326098),
	// 		RootFileSystems: to.Ptr[int64](2),
	// 	},
	// }
}
