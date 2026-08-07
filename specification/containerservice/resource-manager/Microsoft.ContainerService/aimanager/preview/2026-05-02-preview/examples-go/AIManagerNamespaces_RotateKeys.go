package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-05-02-preview/AIManagerNamespaces_RotateKeys.json
func ExampleAIManagerNamespacesClient_RotateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAIManagerNamespacesClient().RotateKeys(ctx, "rgaimanagers", "aimanager1", "namespace-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerserviceaimanager.AIManagerNamespacesClientRotateKeysResponse{
	// 	NamespaceAccessInfo: armcontainerserviceaimanager.NamespaceAccessInfo{
	// 		Endpoint: to.Ptr("https://team-alpha.aks-cluster.eastus.aksapp.io/v1"),
	// 		PrimaryKey: to.Ptr("22222222222222222222222222222222"),
	// 		SecondaryKey: to.Ptr("00000000000000000000000000000000"),
	// 		LastRotatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-23T10:00:00Z"); return t}()),
	// 	},
	// }
}
