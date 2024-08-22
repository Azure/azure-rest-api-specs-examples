package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ConnectedEnvironmentsCertificates_ListByConnectedEnvironment.json
func ExampleConnectedEnvironmentsCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedEnvironmentsCertificatesClient().NewListPager("examplerg", "testcontainerenv", nil)
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
		// page.CertificateCollection = armappcontainers.CertificateCollection{
		// 	Value: []*armappcontainers.Certificate{
		// 		{
		// 			Name: to.Ptr("certificate-firendly-name"),
		// 			Type: to.Ptr("Microsoft.App/ConnectedEnvironments/Certificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/testcontainerenv/certificate-firendly-name"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.CertificateProperties{
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00.000Z"); return t}()),
		// 				IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-06T04:00:00.000Z"); return t}()),
		// 				Issuer: to.Ptr("Issuer Name"),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectAlternativeNames: []*string{
		// 					to.Ptr("CN=my-subject-name.com")},
		// 					SubjectName: to.Ptr("my-subject-name.company.country.net"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 					Valid: to.Ptr(true),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("certificate-firendly-name"),
		// 				Type: to.Ptr("Microsoft.App/ConnectedEnvironments/Certificates"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/testcontainerenv/certificate-firendly-name"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armappcontainers.CertificateProperties{
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00.000Z"); return t}()),
		// 					IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-06T04:00:00.000Z"); return t}()),
		// 					Issuer: to.Ptr("Issuer Name"),
		// 					ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 					SubjectAlternativeNames: []*string{
		// 						to.Ptr("CN=my-subject-name.com")},
		// 						SubjectName: to.Ptr("my-subject-name.company.country.net"),
		// 						Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 						Valid: to.Ptr(true),
		// 					},
		// 			}},
		// 		}
	}
}
