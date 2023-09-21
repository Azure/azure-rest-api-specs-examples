package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listToken.json
func ExampleApplicationsClient_ListTokens() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().ListTokens(ctx, "rg", "myManagedApplication", armmanagedapplications.ListTokenRequest{
		AuthorizationAudience: to.Ptr("authorizationAudience"),
		UserAssignedIdentities: []*string{
			to.Ptr("IdentityOne"),
			to.Ptr("IdentityTwo")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedIdentityTokenResult = armmanagedapplications.ManagedIdentityTokenResult{
	// 	Value: []*armmanagedapplications.ManagedIdentityToken{
	// 		{
	// 			AccessToken: to.Ptr("access_token"),
	// 			AuthorizationAudience: to.Ptr("authorizationAudience"),
	// 			ExpiresIn: to.Ptr("1500"),
	// 			ExpiresOn: to.Ptr("1500"),
	// 			NotBefore: to.Ptr("1500"),
	// 			ResourceID: to.Ptr("resourceId"),
	// 			TokenType: to.Ptr("Bearer"),
	// 	}},
	// }
}
