package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/stable/2024-11-01/examples/Diagnostics_ListAppServiceCertificateOrderDetectorResponse.json
func ExampleCertificateOrdersDiagnosticsClient_NewListAppServiceCertificateOrderDetectorResponsePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateOrdersDiagnosticsClient().NewListAppServiceCertificateOrderDetectorResponsePager("Sample-WestUSResourceGroup", "SampleCertificateOrderName", nil)
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
		// page.DetectorResponseCollection = armappservice.DetectorResponseCollection{
		// 	Value: []*armappservice.DetectorResponse{
		// 		{
		// 			Name: to.Ptr("CertsImport"),
		// 			Type: to.Ptr("Microsoft.Web/certificateOrders/detectors"),
		// 			ID: to.Ptr("/subscriptions/5700fc96-77b4-4f8d-afce-c353d8c443bd/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName/detectors/CertsImport"),
		// 			Properties: &armappservice.DetectorResponseProperties{
		// 				Dataset: []*armappservice.DiagnosticData{
		// 				},
		// 				Metadata: &armappservice.DetectorInfo{
		// 					Name: to.Ptr("Certificates Import/Upload"),
		// 					Author: to.Ptr(""),
		// 					ID: to.Ptr("CertsImport"),
		// 				},
		// 				Status: &armappservice.Status{
		// 					StatusID: to.Ptr(armappservice.InsightStatusNone),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
