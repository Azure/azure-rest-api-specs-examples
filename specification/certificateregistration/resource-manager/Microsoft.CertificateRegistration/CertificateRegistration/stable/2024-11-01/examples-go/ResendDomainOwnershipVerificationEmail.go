package armcertificateregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/certificateregistration/armcertificateregistration"
)

// Generated from example definition: 2024-11-01/ResendDomainOwnershipVerificationEmail.json
func ExampleAppServiceCertificateOrdersClient_ResendRequestEmails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcertificateregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAppServiceCertificateOrdersClient().ResendRequestEmails(ctx, "testrg123", "SampleCertificateOrderName", armcertificateregistration.NameIdentifier{
		Name: to.Ptr("Domain name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
