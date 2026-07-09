package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-03-01-preview/CacheRuleUpdate.json
func ExampleCacheRulesClient_BeginUpdate() {
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
			CredentialSetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/credentialSets/myCredentialSet2"),
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
	// 			CredentialSetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/credentialSets/myCredentialSet2"),
	// 			SourceRepository: to.Ptr("docker.io/library/hello-world"),
	// 			TargetRepository: to.Ptr("cached-docker-hub/hello-world"),
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T18:20:33.8374968+00:00"); return t}()),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
