package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/Diagnostics_ExecuteSiteDetector.json
func ExampleDiagnosticsClient_ExecuteSiteDetector_executeSiteDetector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticsClient().ExecuteSiteDetector(ctx, "Sample-WestUSResourceGroup", "SampleApp", "sitecrashes", "availability", &armappservice.DiagnosticsClientExecuteSiteDetectorOptions{StartTime: nil,
		EndTime:   nil,
		TimeGrain: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticDetectorResponse = armappservice.DiagnosticDetectorResponse{
	// 	Name: to.Ptr("sitecrashes"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/diagnostics/availability/detectors/sitecrashes"),
	// 	Properties: &armappservice.DiagnosticDetectorResponseProperties{
	// 		AbnormalTimePeriods: []*armappservice.DetectorAbnormalTimePeriod{
	// 			{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:50:00.000Z"); return t}()),
	// 				Solutions: []*armappservice.Solution{
	// 				},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-05T22:50:00.000Z"); return t}()),
	// 		}},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:50:00.000Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-05T22:50:00.000Z"); return t}()),
	// 	},
	// }
}
