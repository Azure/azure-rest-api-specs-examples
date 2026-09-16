package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/Subscriptions_GetCustomDomainVerificationId.json
func ExampleContainerAppsAPIClient_GetCustomDomainVerificationID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("d27c3573-f76e-4b26-b871-0ccd2203d08c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsAPIClient().GetCustomDomainVerificationID(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ContainerAppsAPIClientGetCustomDomainVerificationIDResponse{
	// 	Value: to.Ptr("5B406D5E790BBD224468CE0AA814C396203C7CE755F135A80E35D41865E51967"),
	// }
}
