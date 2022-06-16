package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2021-02-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
func ExampleCertificateOrdersDiagnosticsClient_GetAppServiceCertificateOrderDetectorResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewCertificateOrdersDiagnosticsClient("<subscription-id>", cred, nil)
	res, err := client.GetAppServiceCertificateOrderDetectorResponse(ctx,
		"<resource-group-name>",
		"<certificate-order-name>",
		"<detector-name>",
		&armappservice.CertificateOrdersDiagnosticsGetAppServiceCertificateOrderDetectorResponseOptions{StartTime: nil,
			EndTime:   nil,
			TimeGrain: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DetectorResponse.ID: %s\n", *res.ID)
}
