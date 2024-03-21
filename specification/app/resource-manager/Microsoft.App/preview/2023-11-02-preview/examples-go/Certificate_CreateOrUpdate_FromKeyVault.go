package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Certificate_CreateOrUpdate_FromKeyVault.json
func ExampleCertificatesClient_CreateOrUpdate_createOrUpdateCertificateUsingManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().CreateOrUpdate(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", &armappcontainers.CertificatesClientCreateOrUpdateOptions{CertificateEnvelope: &armappcontainers.Certificate{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.CertificateProperties{
			CertificateKeyVaultProperties: &armappcontainers.CertificateKeyVaultProperties{
				Identity:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/microsoft.managedidentity/userassignedidentities/test-user-mi"),
				KeyVaultURL: to.Ptr("https://xxxxxxxx.vault.azure.net/certificates/certName"),
			},
			CertificateType: to.Ptr(armappcontainers.CertificateTypeServerSSLCertificate),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Certificate = armappcontainers.Certificate{
	// 	Type: to.Ptr("Microsoft.App/ManagedEnvironments/Certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificate-firendly-name"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.CertificateProperties{
	// 		CertificateKeyVaultProperties: &armappcontainers.CertificateKeyVaultProperties{
	// 			Identity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/microsoft.managedidentity/userassignedidentities/test-user-mi"),
	// 			KeyVaultURL: to.Ptr("https://xxxxxxxx.vault.azure.net/certificates/certName"),
	// 		},
	// 		CertificateType: to.Ptr(armappcontainers.CertificateTypeServerSSLCertificate),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00.000Z"); return t}()),
	// 		IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-06T04:00:00.000Z"); return t}()),
	// 		Issuer: to.Ptr("Issuer Name"),
	// 		ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
	// 		SubjectAlternativeNames: []*string{
	// 			to.Ptr("CN=my-subject-name.com")},
	// 			SubjectName: to.Ptr("my-subject-name.company.country.net"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 			Valid: to.Ptr(true),
	// 		},
	// 	}
}
