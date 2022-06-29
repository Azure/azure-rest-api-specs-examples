package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/FederatedIdentityCredentialCreate.json
func ExampleFederatedIdentityCredentialsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewFederatedIdentityCredentialsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rgName",
		"resourceName",
		"ficResourceName",
		armmsi.FederatedIdentityCredential{
			Properties: &armmsi.FederatedIdentityCredentialProperties{
				Audiences: []*string{
					to.Ptr("api://AzureADTokenExchange")},
				Issuer:  to.Ptr("https://oidc.prod-aks.azure.com/IssuerGUID"),
				Subject: to.Ptr("system:serviceaccount:ns:svcaccount"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
