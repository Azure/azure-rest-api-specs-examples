package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/UsageMetrics_ListByWorkspace_MaximumSet_Gen.json
func ExampleUsageMetricsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("5C707B5F-6130-4F71-819E-953A28942E88", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsageMetricsClient().NewListByWorkspacePager("rgiotfirmwaredefense", "exampleWorkspaceName", nil)
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
		// page = armiotfirmwaredefense.UsageMetricsClientListByWorkspaceResponse{
		// 	UsageMetricListResult: armiotfirmwaredefense.UsageMetricListResult{
		// 		Value: []*armiotfirmwaredefense.UsageMetric{
		// 			{
		// 				Properties: &armiotfirmwaredefense.UsageMetricProperties{
		// 					MonthlyFirmwareUploadCount: to.Ptr[int64](26),
		// 					TotalFirmwareCount: to.Ptr[int64](28),
		// 					ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/exampleWorkspaceName/usageMetrics/current"),
		// 				Name: to.Ptr("apalfbnzccynphwqeyzk"),
		// 				Type: to.Ptr("jtvcvsrlngowwoapmmdeevho"),
		// 				SystemData: &armiotfirmwaredefense.SystemData{
		// 					CreatedBy: to.Ptr("nqisshvdzqcxzbujvacin"),
		// 					CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("of"),
		// 					LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/avb"),
		// 	},
		// }
	}
}
