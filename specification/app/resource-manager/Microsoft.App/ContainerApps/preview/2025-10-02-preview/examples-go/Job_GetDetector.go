package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/Job_GetDetector.json
func ExampleJobsClient_GetDetector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("f07f3711-b45e-40fe-a941-4e6d93f851e6", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().GetDetector(ctx, "mikono-workerapp-test-rg", "mikonojob1", "containerappjobnetworkIO", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.JobsClientGetDetectorResponse{
	// 	Diagnostics: armappcontainers.Diagnostics{
	// 		ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/jobs/mikonojob1/detectors/containerappjobnetworkIO"),
	// 		Name: to.Ptr("containerappjobnetworkIO"),
	// 		Type: to.Ptr("Microsoft.App/jobs/detectors"),
	// 		Properties: &armappcontainers.DiagnosticsProperties{
	// 			Metadata: &armappcontainers.DiagnosticsDefinition{
	// 				ID: to.Ptr("containerappjobnetworkIO"),
	// 				Name: to.Ptr("Container App Job Network Inbound and Outbound"),
	// 				Description: to.Ptr("This detector shows the Container App Job Network Inbound and Outbound."),
	// 				Author: to.Ptr(""),
	// 				Category: to.Ptr("Availability and Performance"),
	// 				SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
	// 				},
	// 				Type: to.Ptr("Detector"),
	// 				Score: to.Ptr[float32](0),
	// 			},
	// 			Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
	// 				{
	// 					Table: &armappcontainers.DiagnosticDataTableResponseObject{
	// 						TableName: to.Ptr(""),
	// 						Columns: []*armappcontainers.DiagnosticDataTableResponseColumn{
	// 							{
	// 								ColumnName: to.Ptr("TimeStamp"),
	// 								DataType: to.Ptr("DateTime"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Metric"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Average"),
	// 								DataType: to.Ptr("Double"),
	// 							},
	// 						},
	// 						Rows: []any{
	// 							[]any{
	// 								"2022-03-15T21:35:00",
	// 								"RxBytes",
	// 								0,
	// 							},
	// 						},
	// 					},
	// 					RenderingProperties: &armappcontainers.DiagnosticRendering{
	// 						Type: to.Ptr[int32](8),
	// 						Title: to.Ptr("Container App Job Network Inbound "),
	// 						Description: to.Ptr(""),
	// 						IsVisible: to.Ptr(true),
	// 					},
	// 				},
	// 			},
	// 			Status: &armappcontainers.DiagnosticsStatus{
	// 				StatusID: to.Ptr[int32](3),
	// 			},
	// 		},
	// 	},
	// }
}
