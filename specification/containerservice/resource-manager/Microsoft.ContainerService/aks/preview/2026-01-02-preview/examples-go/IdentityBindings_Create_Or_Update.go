package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-01-02-preview/IdentityBindings_Create_Or_Update.json
func ExampleIdentityBindingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIdentityBindingsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "identitybinding1", armcontainerservice.IdentityBinding{
		Properties: &armcontainerservice.IdentityBindingProperties{
			ManagedIdentity: &armcontainerservice.IdentityBindingManagedIdentityProfile{
				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
			},
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
	// res = armcontainerservice.IdentityBindingsClientCreateOrUpdateResponse{
	// 	IdentityBinding: &armcontainerservice.IdentityBinding{
	// 		Name: to.Ptr("identitybinding1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/identityBindings"),
	// 		ETag: to.Ptr("string"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/identityBindings/identitybinding1"),
	// 		Properties: &armcontainerservice.IdentityBindingProperties{
	// 			ManagedIdentity: &armcontainerservice.IdentityBindingManagedIdentityProfile{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			OidcIssuer: &armcontainerservice.IdentityBindingOidcIssuerProfile{
	// 				OidcIssuerURL: to.Ptr("https://oidc-endpoint"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerservice.IdentityBindingProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
