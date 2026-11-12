package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v4"
)

// Generated from example definition: 2026-10-15-preview/CacheRuleUpdateEntraAuthentication.json
func ExampleCacheRulesClient_BeginUpdate_updateEntraAuthenticationWhilePreservingOmittedFields() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCacheRulesClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myCacheRule", armcontainerregistry.CacheRuleUpdateParameters{
		Properties: &armcontainerregistry.CacheRuleUpdateProperties{
			AdditionalAuthenticationProperties: &armcontainerregistry.EntraFederatedIdentityCredentialAuthenticationUpdateProperties{
				AuthenticationType:           to.Ptr(armcontainerregistry.AdditionalAuthenticationTypeEntraFederatedIdentityCredential),
				FederatedApplicationClientID: to.Ptr("fedcba98-7654-3210-abcd-ef1234567890"),
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
	// res = armcontainerregistry.CacheRulesClientUpdateResponse{
	// 	CacheRule: armcontainerregistry.CacheRule{
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/cacheRules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/cacheRules/myCacheRule"),
	// 		Name: to.Ptr("myCacheRule"),
	// 		Properties: &armcontainerregistry.CacheRuleProperties{
	// 			SourceRepository: to.Ptr("myupstream.azurecr.io/hello-world"),
	// 			TargetRepository: to.Ptr("cached-upstream/hello-world"),
	// 			AdditionalAuthenticationProperties: &armcontainerregistry.EntraFederatedIdentityCredentialAuthenticationProperties{
	// 				AuthenticationType: to.Ptr(armcontainerregistry.AdditionalAuthenticationTypeEntraFederatedIdentityCredential),
	// 				SourceTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				FederatedApplicationClientID: to.Ptr("fedcba98-7654-3210-abcd-ef1234567890"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
