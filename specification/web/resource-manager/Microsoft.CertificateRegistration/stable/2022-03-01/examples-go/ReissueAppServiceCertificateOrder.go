package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/ReissueAppServiceCertificateOrder.json
func ExampleCertificateOrdersClient_Reissue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewCertificateOrdersClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Reissue(ctx,
		"testrg123",
		"SampleCertificateOrderName",
		armappservice.ReissueCertificateOrderRequest{
			Properties: &armappservice.ReissueCertificateOrderRequestProperties{
				Csr:                        to.Ptr("CSR1223238Value"),
				DelayExistingRevokeInHours: to.Ptr[int32](2),
				IsPrivateKeyExternal:       to.Ptr(false),
				KeySize:                    to.Ptr[int32](2048),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
