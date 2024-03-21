package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Builders_ListBySubscription.json
func ExampleBuildersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBuildersClient().NewListBySubscriptionPager(nil)
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
		// page.BuilderCollection = armappcontainers.BuilderCollection{
		// 	Value: []*armappcontainers.BuilderResource{
		// 		{
		// 			Name: to.Ptr("testBuilder1"),
		// 			Type: to.Ptr("Microsoft.App/builders"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.App/builders/testBuilder1"),
		// 			SystemData: &armappcontainers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample@microsoft.com"),
		// 				CreatedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Identity: &armappcontainers.ManagedServiceIdentity{
		// 				Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armappcontainers.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armappcontainers.BuilderProperties{
		// 				ContainerRegistries: []*armappcontainers.ContainerRegistry{
		// 					{
		// 						ContainerRegistryServer: to.Ptr("test.azurecr.io"),
		// 						IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
		// 					},
		// 					{
		// 						ContainerRegistryServer: to.Ptr("test2.azurecr.io"),
		// 						IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
		// 				}},
		// 				EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/managedEnvironments/testEnv"),
		// 				ProvisioningState: to.Ptr(armappcontainers.BuilderProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testBuilder2"),
		// 			Type: to.Ptr("Microsoft.App/builders"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg2/providers/Microsoft.App/builders/testBuilder2"),
		// 			SystemData: &armappcontainers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample@microsoft.com"),
		// 				CreatedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Identity: &armappcontainers.ManagedServiceIdentity{
		// 				Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armappcontainers.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armappcontainers.BuilderProperties{
		// 				ContainerRegistries: []*armappcontainers.ContainerRegistry{
		// 					{
		// 						ContainerRegistryServer: to.Ptr("test.azurecr.io"),
		// 						IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
		// 					},
		// 					{
		// 						ContainerRegistryServer: to.Ptr("test2.azurecr.io"),
		// 						IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
		// 				}},
		// 				EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/managedEnvironments/testEnv"),
		// 				ProvisioningState: to.Ptr(armappcontainers.BuilderProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
