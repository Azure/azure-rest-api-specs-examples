package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-01-01/examples/RenewAppServiceCertificateOrder.json
func ExampleCertificateOrdersClient_Renew() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificateOrdersClient().Renew(ctx, "testrg123", "SampleCertificateOrderName", armappservice.RenewCertificateOrderRequest{
		Properties: &armappservice.RenewCertificateOrderRequestProperties{
			Csr:                  to.Ptr("CSR1223238Value"),
			IsPrivateKeyExternal: to.Ptr(false),
			KeySize:              to.Ptr[int32](2048),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
