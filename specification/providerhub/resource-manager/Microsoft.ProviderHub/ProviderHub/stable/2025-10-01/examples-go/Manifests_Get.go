package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/Manifests_Get.json
func ExampleManifestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManifestsClient().Get(ctx, "Microsoft.Contoso", "prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.ManifestsClientGetResponse{
	// 	ManifestInfo: armproviderhub.ManifestInfo{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/manifests/prod"),
	// 		Name: to.Ptr("prod"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/manifest"),
	// 		Properties: &armproviderhub.ManifestInfoProperties{
	// 			Manifest: to.Ptr("<<Core RP manifest>>"),
	// 			CommitID: to.Ptr("CD2805AE-8017-47DD-B720-CF6CBF28A9D7"),
	// 		},
	// 	},
	// }
}
