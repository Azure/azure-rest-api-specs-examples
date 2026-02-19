package armartifactsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/artifactsigning/armartifactsigning"
)

// Generated from example definition: 2025-10-13/CodeSigningAccounts_CheckNameAvailability.json
func ExampleCodeSigningAccountsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armartifactsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCodeSigningAccountsClient().CheckNameAvailability(ctx, armartifactsigning.CheckNameAvailability{
		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
		Name: to.Ptr("sample-account"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armartifactsigning.CodeSigningAccountsClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResult: &armartifactsigning.CheckNameAvailabilityResult{
	// 		NameAvailable: to.Ptr(true),
	// 	},
	// }
}
