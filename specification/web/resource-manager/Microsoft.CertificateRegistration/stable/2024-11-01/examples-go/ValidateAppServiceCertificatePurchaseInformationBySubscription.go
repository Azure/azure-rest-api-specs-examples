package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2024-11-01/examples/ValidateAppServiceCertificatePurchaseInformationBySubscription.json
func ExampleCertificateOrdersClient_ValidatePurchaseInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificateOrdersClient().ValidatePurchaseInformation(ctx, armappservice.CertificateOrder{
		Location: to.Ptr("Global"),
		Properties: &armappservice.CertificateOrderProperties{
			AutoRenew: to.Ptr(true),
			Certificates: map[string]*armappservice.Certificate{
				"SampleCertName1": {
					KeyVaultID:         to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
					KeyVaultSecretName: to.Ptr("SampleSecretName1"),
				},
				"SampleCertName2": {
					KeyVaultID:         to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
					KeyVaultSecretName: to.Ptr("SampleSecretName2"),
				},
			},
			DistinguishedName: to.Ptr("CN=SampleCustomDomain.com"),
			KeySize:           to.Ptr[int32](2048),
			ProductType:       to.Ptr(armappservice.CertificateProductTypeStandardDomainValidatedSSL),
			ValidityInYears:   to.Ptr[int32](2),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
