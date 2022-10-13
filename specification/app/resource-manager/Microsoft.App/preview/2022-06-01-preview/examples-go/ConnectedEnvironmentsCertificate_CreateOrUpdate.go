package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsCertificate_CreateOrUpdate.json
func ExampleConnectedEnvironmentsCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewConnectedEnvironmentsCertificatesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", &armappcontainers.ConnectedEnvironmentsCertificatesClientCreateOrUpdateOptions{CertificateEnvelope: &armappcontainers.Certificate{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.CertificateProperties{
			Password: to.Ptr("private key password"),
			Value:    []byte("PFX-or-PEM-blob"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
