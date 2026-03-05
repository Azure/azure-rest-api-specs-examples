package armcertificateregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/certificateregistration/armcertificateregistration"
)

// Generated from example definition: 2024-11-01/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
func ExampleCertificateOrdersDiagnosticsClient_GetAppServiceCertificateOrderDetectorResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcertificateregistration.NewClientFactory("5700fc96-77b4-4f8d-afce-c353d8c443bd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateOrdersDiagnosticsClient().GetAppServiceCertificateOrderDetectorResponse(ctx, "Sample-WestUSResourceGroup", "SampleCertificateOrderName", "AutoRenewStatus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcertificateregistration.CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseResponse{
	// 	DetectorResponse: &armcertificateregistration.DetectorResponse{
	// 		Name: to.Ptr("AutoRenewStatus"),
	// 		ID: to.Ptr("/subscriptions/5700fc96-77b4-4f8d-afce-c353d8c443bd/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName/detectors/AutoRenewStatus"),
	// 		Properties: &armcertificateregistration.DetectorResponseProperties{
	// 			Dataset: []*armcertificateregistration.DiagnosticData{
	// 				{
	// 					RenderingProperties: &armcertificateregistration.Rendering{
	// 						Type: to.Ptr(armcertificateregistration.RenderingTypeInsights),
	// 					},
	// 					Table: &armcertificateregistration.DataTableResponseObject{
	// 						Columns: []*armcertificateregistration.DataTableResponseColumn{
	// 							{
	// 								ColumnName: to.Ptr("Status"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Message"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Data.Name"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Data.Value"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Expanded"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 							{
	// 								ColumnName: to.Ptr("Solutions"),
	// 								DataType: to.Ptr("String"),
	// 							},
	// 						},
	// 						Rows: [][]*string{
	// 							[]*string{
	// 								to.Ptr("Info"),
	// 								to.Ptr("App Service Certificate name SampleCertificateOrderName"),
	// 								to.Ptr("Certificate for"),
	// 								to.Ptr("<b>*.mysampledomain.com</b>"),
	// 								to.Ptr("False"),
	// 								to.Ptr("null"),
	// 							},
	// 							[]*string{
	// 								to.Ptr("Info"),
	// 								to.Ptr("App Service Certificate name SampleCertificateOrderName"),
	// 								to.Ptr("Certificate Status"),
	// 								to.Ptr("<b>Issued</b>"),
	// 								to.Ptr("False"),
	// 								to.Ptr("null"),
	// 							},
	// 						},
	// 						TableName: to.Ptr(""),
	// 					},
	// 				},
	// 			},
	// 			Metadata: &armcertificateregistration.DetectorInfo{
	// 				Name: to.Ptr("AutoRenewStatus"),
	// 				Type: to.Ptr(armcertificateregistration.DetectorTypeDetector),
	// 				Description: to.Ptr("Check auto renew status"),
	// 				ID: to.Ptr("AutoRenewStatus"),
	// 				Score: to.Ptr[float32](0),
	// 				SupportTopicList: []*armcertificateregistration.SupportTopic{
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
