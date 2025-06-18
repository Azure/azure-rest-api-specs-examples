package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2024-11-01/examples/ReissueAppServiceCertificateOrder.json
func ExampleCertificateOrdersClient_Reissue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificateOrdersClient().Reissue(ctx, "testrg123", "SampleCertificateOrderName", armappservice.ReissueCertificateOrderRequest{
		Properties: &armappservice.ReissueCertificateOrderRequestProperties{
			Csr:                        to.Ptr("CSR1223238Value"),
			DelayExistingRevokeInHours: to.Ptr[int32](2),
			IsPrivateKeyExternal:       to.Ptr(false),
			KeySize:                    to.Ptr[int32](2048),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
