package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3855ffb4be0cd4d227b130b67d874fa816736c04/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-08-02-preview/examples/IdentityBindings_Get.json
func ExampleIdentityBindingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIdentityBindingsClient().Get(ctx, "rg1", "clustername1", "identitybinding1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IdentityBinding = armcontainerservice.IdentityBinding{
	// 	Name: to.Ptr("identitybinding1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/identityBindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/identityBindings/identitybinding1"),
	// 	ETag: to.Ptr("string"),
	// 	Properties: &armcontainerservice.IdentityBindingProperties{
	// 		ManagedIdentity: &armcontainerservice.IdentityBindingManagedIdentityProfile{
	// 			ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		OidcIssuer: &armcontainerservice.IdentityBindingOidcIssuerProfile{
	// 			OidcIssuerURL: to.Ptr("https://oidc-endpoint"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservice.IdentityBindingProvisioningStateSucceeded),
	// 	},
	// }
}
