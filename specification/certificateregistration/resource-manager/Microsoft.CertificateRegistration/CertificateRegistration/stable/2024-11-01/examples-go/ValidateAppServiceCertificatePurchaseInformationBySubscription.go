package armcertificateregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/certificateregistration/armcertificateregistration"
)

// Generated from example definition: 2024-11-01/ValidateAppServiceCertificatePurchaseInformationBySubscription.json
func ExampleAppServiceCertificateOrdersClient_ValidatePurchaseInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcertificateregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAppServiceCertificateOrdersClient().ValidatePurchaseInformation(ctx, armcertificateregistration.AppServiceCertificateOrder{
		Location: to.Ptr("Global"),
		Properties: &armcertificateregistration.AppServiceCertificateOrderProperties{
			AutoRenew: to.Ptr(true),
			Certificates: map[string]*armcertificateregistration.AppServiceCertificate{
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
			ProductType:       to.Ptr(armcertificateregistration.CertificateProductTypeStandardDomainValidatedSSL),
			ValidityInYears:   to.Ptr[int32](2),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
