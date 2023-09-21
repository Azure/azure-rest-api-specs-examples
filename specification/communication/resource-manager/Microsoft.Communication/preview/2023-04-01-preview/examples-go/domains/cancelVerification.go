package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d7b535d1273b18623ca0d63a6ebb0456dab95ba/specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/domains/cancelVerification.json
func ExampleDomainsClient_BeginCancelVerification() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDomainsClient().BeginCancelVerification(ctx, "MyResourceGroup", "MyEmailServiceResource", "mydomain.com", armcommunication.VerificationParameter{
		VerificationType: to.Ptr(armcommunication.VerificationTypeSPF),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
