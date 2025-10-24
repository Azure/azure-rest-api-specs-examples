package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b2965096067d6f8374b5485b0568fd36e7c9d099/specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ConnectedEnvironmentsCertificates_Patch.json
func ExampleConnectedEnvironmentsCertificatesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedEnvironmentsCertificatesClient().BeginUpdate(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", armappcontainers.CertificatePatch{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Certificate = armappcontainers.Certificate{
	// 	Type: to.Ptr("Microsoft.App/ConnectedEnvironments/Certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/testcontainerenv/certificates/certificate-firendly-name"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armappcontainers.CertificateProperties{
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
