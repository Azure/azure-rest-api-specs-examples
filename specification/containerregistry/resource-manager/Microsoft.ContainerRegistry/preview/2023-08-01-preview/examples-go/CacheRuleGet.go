package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/CacheRuleGet.json
func ExampleCacheRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCacheRulesClient().Get(ctx, "myResourceGroup", "myRegistry", "myCacheRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CacheRule = armcontainerregistry.CacheRule{
	// 	Name: to.Ptr("myCacheRule"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/cacheRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/cacheRules/myCacheRule"),
	// 	Properties: &armcontainerregistry.CacheRuleProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T18:20:33.8374968+00:00"); return t}()),
	// 		CredentialSetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/credentialSets/myCredentialSet"),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		SourceRepository: to.Ptr("docker.io/library/hello-world"),
	// 		TargetRepository: to.Ptr("cached-docker-hub/hello-world"),
	// 	},
	// }
}
