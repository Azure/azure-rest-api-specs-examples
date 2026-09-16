package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/Certificate_CreateOrUpdate.json
func ExampleCertificatesClient_CreateOrUpdate_createOrUpdateCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().CreateOrUpdate(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", armappcontainers.Certificate{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.CertificateProperties{
			Password: to.Ptr("private key password"),
			Value:    []byte("Y2VydA=="),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.CertificatesClientCreateOrUpdateResponse{
	// 	Certificate: armappcontainers.Certificate{
	// 		Type: to.Ptr("Microsoft.App/ManagedEnvironments/Certificates"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-firendly-name"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappcontainers.CertificateProperties{
	// 			ExpirationDate: to.Ptr(time.Date(2022, time.November, 6, 4, 0, 0, 0, time.UTC)),
	// 			IssueDate: to.Ptr(time.Date(2021, time.November, 6, 4, 0, 0, 0, time.UTC)),
	// 			Issuer: to.Ptr("Issuer Name"),
	// 			ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
	// 			SubjectAlternativeNames: []*string{
	// 				to.Ptr("CN=my-subject-name.com"),
	// 			},
	// 			SubjectName: to.Ptr("my-subject-name.company.country.net"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 			Valid: to.Ptr(true),
	// 		},
	// 	},
	// }
}
