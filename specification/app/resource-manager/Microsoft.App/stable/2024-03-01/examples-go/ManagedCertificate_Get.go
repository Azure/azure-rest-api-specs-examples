package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificate_Get.json
func ExampleManagedCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedCertificatesClient().Get(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedCertificate = armappcontainers.ManagedCertificate{
	// 	Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ManagedCertificateProperties{
	// 		DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
	// 		ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
	// 		SubjectName: to.Ptr("CN=my-subject-name.company.country.net"),
	// 	},
	// }
}
