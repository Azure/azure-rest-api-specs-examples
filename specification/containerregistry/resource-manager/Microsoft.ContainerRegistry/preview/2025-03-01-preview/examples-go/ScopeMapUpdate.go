package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ScopeMapUpdate.json
func ExampleScopeMapsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScopeMapsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myScopeMap", armcontainerregistry.ScopeMapUpdateParameters{
		Properties: &armcontainerregistry.ScopeMapPropertiesUpdateParameters{
			Description: to.Ptr("Developer Scopes"),
			Actions: []*string{
				to.Ptr("repositories/myrepository/contentWrite"),
				to.Ptr("repositories/myrepository/contentRead")},
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
	// res.ScopeMap = armcontainerregistry.ScopeMap{
	// 	Name: to.Ptr("myScopeMap"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/scopeMaps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myScopeMap"),
	// 	Properties: &armcontainerregistry.ScopeMapProperties{
	// 		Type: to.Ptr("IsUserDefined"),
	// 		Actions: []*string{
	// 			to.Ptr("repositories/myrepository/contentWrite"),
	// 			to.Ptr("repositories/myrepository/contentRead")},
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.070Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
