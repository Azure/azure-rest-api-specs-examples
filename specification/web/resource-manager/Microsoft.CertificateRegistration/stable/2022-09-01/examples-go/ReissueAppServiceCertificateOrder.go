package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-09-01/examples/ReissueAppServiceCertificateOrder.json
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
