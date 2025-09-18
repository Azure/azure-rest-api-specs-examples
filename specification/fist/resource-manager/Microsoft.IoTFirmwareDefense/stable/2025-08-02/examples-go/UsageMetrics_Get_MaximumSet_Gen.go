package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-08-02/UsageMetrics_Get_MaximumSet_Gen.json
func ExampleUsageMetricsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUsageMetricsClient().Get(ctx, "rgiotfirmwaredefense", "exampleWorkspaceName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotfirmwaredefense.UsageMetricsClientGetResponse{
	// 	UsageMetric: &armiotfirmwaredefense.UsageMetric{
	// 		Properties: &armiotfirmwaredefense.UsageMetricProperties{
	// 			MonthlyFirmwareUploadCount: to.Ptr[int64](25),
	// 			TotalFirmwareCount: to.Ptr[int64](29),
	// 			ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/exampleWorkspaceName"),
	// 		Name: to.Ptr("usageMetrics"),
	// 		Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/usageMetrics"),
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
