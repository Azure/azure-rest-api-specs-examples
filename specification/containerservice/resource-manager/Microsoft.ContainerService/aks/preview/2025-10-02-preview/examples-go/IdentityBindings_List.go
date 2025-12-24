package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/IdentityBindings_List.json
func ExampleIdentityBindingsClient_NewListByManagedClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIdentityBindingsClient().NewListByManagedClusterPager("rg1", "clustername1", nil)
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
		// page.IdentityBindingListResult = armcontainerservice.IdentityBindingListResult{
		// 	Value: []*armcontainerservice.IdentityBinding{
		// 		{
		// 			Name: to.Ptr("identitybinding1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/identityBindings"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/identityBindings/identitybinding1"),
		// 			ETag: to.Ptr("string"),
		// 			Properties: &armcontainerservice.IdentityBindingProperties{
		// 				ManagedIdentity: &armcontainerservice.IdentityBindingManagedIdentityProfile{
		// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 				OidcIssuer: &armcontainerservice.IdentityBindingOidcIssuerProfile{
		// 					OidcIssuerURL: to.Ptr("https://oidc-endpoint"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerservice.IdentityBindingProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
