package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v4"
)

// Generated from example definition: 2026-10-15-preview/CacheRuleCreateUserAssignedMIEntraFederatedIdentityCredential.json
func ExampleCacheRulesClient_BeginCreate_cacheRuleCreateUserAssignedMiAuthEntraFederatedIdentityCredential() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCacheRulesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myCacheRule", armcontainerregistry.CacheRule{
		Identity: &armcontainerregistry.IdentityProperties{
			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
				"/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity": {},
			},
		},
		Properties: &armcontainerregistry.CacheRuleProperties{
			SourceRepository: to.Ptr("myupstream.azurecr.io/hello-world"),
			TargetRepository: to.Ptr("cached-upstream/hello-world"),
			AdditionalAuthenticationProperties: &armcontainerregistry.EntraFederatedIdentityCredentialAuthenticationProperties{
				AuthenticationType:           to.Ptr(armcontainerregistry.AdditionalAuthenticationTypeEntraFederatedIdentityCredential),
				SourceTenantID:               to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
				FederatedApplicationClientID: to.Ptr("abcdef12-3456-7890-abcd-ef1234567890"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.CacheRulesClientCreateResponse{
	// 	CacheRule: armcontainerregistry.CacheRule{
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/cacheRules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/cacheRules/myCacheRule"),
	// 		Name: to.Ptr("myCacheRule"),
	// 		Identity: &armcontainerregistry.IdentityProperties{
	// 			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
	// 				"/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity": &armcontainerregistry.UserIdentityProperties{
	// 					ClientID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 					PrincipalID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armcontainerregistry.CacheRuleProperties{
	// 			SourceRepository: to.Ptr("myupstream.azurecr.io/hello-world"),
	// 			TargetRepository: to.Ptr("cached-upstream/hello-world"),
	// 			AdditionalAuthenticationProperties: &armcontainerregistry.EntraFederatedIdentityCredentialAuthenticationProperties{
	// 				AuthenticationType: to.Ptr(armcontainerregistry.AdditionalAuthenticationTypeEntraFederatedIdentityCredential),
	// 				SourceTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				FederatedApplicationClientID: to.Ptr("abcdef12-3456-7890-abcd-ef1234567890"),
	// 			},
	// 			CreationDate: to.Ptr(time.Date(2022, time.November, 7, 18, 20, 33, 837496800, time.UTC)),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
