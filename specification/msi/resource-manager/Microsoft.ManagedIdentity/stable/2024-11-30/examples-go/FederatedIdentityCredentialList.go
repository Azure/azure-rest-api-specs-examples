package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/FederatedIdentityCredentialList.json
func ExampleFederatedIdentityCredentialsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmsi.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFederatedIdentityCredentialsClient().NewListPager("rgName", "resourceName", &armmsi.FederatedIdentityCredentialsClientListOptions{Top: nil,
		Skiptoken: nil,
	})
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
		// page.FederatedIdentityCredentialsListResult = armmsi.FederatedIdentityCredentialsListResult{
		// 	Value: []*armmsi.FederatedIdentityCredential{
		// 		{
		// 			Name: to.Ptr("ficResourceName"),
		// 			Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"),
		// 			ID: to.Ptr("/subscriptions/c267c0e7-0a73-4789-9e17-d26aeb0904e5/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName/federatedIdentityCredentials/ficResourceName"),
		// 			Properties: &armmsi.FederatedIdentityCredentialProperties{
		// 				Audiences: []*string{
		// 					to.Ptr("api://AzureADTokenExchange")},
		// 					Issuer: to.Ptr("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID"),
		// 					Subject: to.Ptr("system:serviceaccount:ns:svcaccount"),
		// 				},
		// 		}},
		// 	}
	}
}
