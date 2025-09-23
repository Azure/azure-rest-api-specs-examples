package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/Builders_CreateOrUpdate.json
func ExampleBuildersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBuildersClient().BeginCreateOrUpdate(ctx, "rg", "testBuilder", armappcontainers.BuilderResource{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"company": to.Ptr("Microsoft"),
		},
		Identity: &armappcontainers.ManagedServiceIdentity{
			Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
		Properties: &armappcontainers.BuilderProperties{
			ContainerRegistries: []*armappcontainers.ContainerRegistry{
				{
					ContainerRegistryServer: to.Ptr("test.azurecr.io"),
					IdentityResourceID:      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
				},
				{
					ContainerRegistryServer: to.Ptr("test2.azurecr.io"),
					IdentityResourceID:      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
				}},
			EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/managedEnvironments/testEnv"),
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
	// res.BuilderResource = armappcontainers.BuilderResource{
	// 	Name: to.Ptr("testBuilder"),
	// 	Type: to.Ptr("Microsoft.App/builders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/builders/testBuilder"),
	// 	SystemData: &armappcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample@microsoft.com"),
	// 		CreatedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.494Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"company": to.Ptr("Microsoft"),
	// 	},
	// 	Identity: &armappcontainers.ManagedServiceIdentity{
	// 		Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armappcontainers.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armappcontainers.BuilderProperties{
	// 		ContainerRegistries: []*armappcontainers.ContainerRegistry{
	// 			{
	// 				ContainerRegistryServer: to.Ptr("test.azurecr.io"),
	// 				IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 			},
	// 			{
	// 				ContainerRegistryServer: to.Ptr("test2.azurecr.io"),
	// 				IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 		}},
	// 		EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/managedEnvironments/testEnv"),
	// 		ProvisioningState: to.Ptr(armappcontainers.BuilderProvisioningStateSucceeded),
	// 	},
	// }
}
