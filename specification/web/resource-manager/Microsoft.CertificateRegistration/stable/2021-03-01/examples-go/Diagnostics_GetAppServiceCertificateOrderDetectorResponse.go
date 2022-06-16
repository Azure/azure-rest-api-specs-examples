package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2021-03-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
func ExampleCertificateOrdersDiagnosticsClient_GetAppServiceCertificateOrderDetectorResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewCertificateOrdersDiagnosticsClient("5700fc96-77b4-4f8d-afce-c353d8c443bd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAppServiceCertificateOrderDetectorResponse(ctx,
		"Sample-WestUSResourceGroup",
		"SampleCertificateOrderName",
		"AutoRenewStatus",
		&armappservice.CertificateOrdersDiagnosticsClientGetAppServiceCertificateOrderDetectorResponseOptions{StartTime: nil,
			EndTime:   nil,
			TimeGrain: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
