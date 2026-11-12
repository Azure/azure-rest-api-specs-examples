package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v4"
)

// Generated from example definition: 2026-10-15-preview/CacheRuleUpdateGoogleAuthentication.json
func ExampleCacheRulesClient_BeginUpdate_updateGoogleAuthenticationWhilePreservingOmittedFields() {
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
			AdditionalAuthenticationProperties: &armcontainerregistry.GARAuthenticationUpdateProperties{
				AuthenticationType:       to.Ptr(armcontainerregistry.AdditionalAuthenticationTypeGoogleArtifactRegistry),
				WorkloadIdentityProvider: to.Ptr("updated-identity-provider"),
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
	// 			SourceRepository: to.Ptr("us-west1-docker.pkg.dev/my-project/my-repository/hello-world"),
	// 			TargetRepository: to.Ptr("cached-gar/hello-world"),
	// 			AdditionalAuthenticationProperties: &armcontainerregistry.GarAuthenticationProperties{
	// 				AuthenticationType: to.Ptr(armcontainerregistry.AdditionalAuthenticationTypeGoogleArtifactRegistry),
	// 				ProjectNumber: to.Ptr("123456789012"),
	// 				WorkloadIdentityPool: to.Ptr("my-workload-identity-pool"),
	// 				WorkloadIdentityProvider: to.Ptr("updated-identity-provider"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
