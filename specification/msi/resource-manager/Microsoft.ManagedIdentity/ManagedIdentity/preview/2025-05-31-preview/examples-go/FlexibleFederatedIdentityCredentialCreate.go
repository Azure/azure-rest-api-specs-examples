package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: 2025-05-31-preview/FlexibleFederatedIdentityCredentialCreate.json
func ExampleFederatedIdentityCredentialsClient_CreateOrUpdate_flexibleFederatedIdentityCredentialCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmsi.NewClientFactory("c267c0e7-0a73-4789-9e17-d26aeb0904e5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFederatedIdentityCredentialsClient().CreateOrUpdate(ctx, "rgName", "resourceName", "ficResourceName", armmsi.FederatedIdentityCredential{
		Properties: &armmsi.FederatedIdentityCredentialProperties{
			Issuer: to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
			Audiences: []*string{
				to.Ptr("api://AzureADTokenExchange"),
			},
			ClaimsMatchingExpression: &armmsi.ClaimsMatchingExpression{
				LanguageVersion: to.Ptr[int32](1),
				Value:           to.Ptr("claims['sub'] matches system:serviceaccount:ns:*"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmsi.FederatedIdentityCredentialsClientCreateOrUpdateResponse{
	// 	FederatedIdentityCredential: armmsi.FederatedIdentityCredential{
	// 		ID: to.Ptr("/subscriptions/c267c0e7-0a73-4789-9e17-d26aeb0904e5/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName/federatedIdentityCredentials/ficResourceName"),
	// 		Name: to.Ptr("ficResourceName"),
	// 		Properties: &armmsi.FederatedIdentityCredentialProperties{
	// 			Issuer: to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
	// 			Audiences: []*string{
	// 				to.Ptr("api://AzureADTokenExchange"),
	// 			},
	// 			ClaimsMatchingExpression: &armmsi.ClaimsMatchingExpression{
	// 				LanguageVersion: to.Ptr[int32](1),
	// 				Value: to.Ptr("claims['sub'] matches system:serviceaccount:ns:*"),
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"),
	// 	},
	// }
}
