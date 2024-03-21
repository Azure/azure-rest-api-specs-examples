package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ManagedCertificates_ListByManagedEnvironment.json
func ExampleManagedCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedCertificatesClient().NewListPager("examplerg", "testcontainerenv", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ManagedCertificateCollection = armappcontainers.ManagedCertificateCollection{
		// 	Value: []*armappcontainers.ManagedCertificate{
		// 		{
		// 			Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ManagedCertificateProperties{
		// 				DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectName: to.Ptr("CN=my-subject-name.company.country.net"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name-root"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ManagedCertificateProperties{
		// 				DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationHTTP),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectName: to.Ptr("CN=company.country.net"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name-txt"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ManagedCertificateProperties{
		// 				DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationTXT),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectName: to.Ptr("CN=txt.company.country.net"),
		// 			},
		// 	}},
		// }
	}
}
