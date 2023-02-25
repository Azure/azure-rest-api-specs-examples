package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ScopeMapList.json
func ExampleScopeMapsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewScopeMapsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myResourceGroup", "myRegistry", nil)
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
		// page.ScopeMapListResult = armcontainerregistry.ScopeMapListResult{
		// 	Value: []*armcontainerregistry.ScopeMap{
		// 		{
		// 			Name: to.Ptr("myScopeMap"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/scopeMaps"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myScopeMap"),
		// 			Properties: &armcontainerregistry.ScopeMapProperties{
		// 				Type: to.Ptr("IsUserDefined"),
		// 				Actions: []*string{
		// 					to.Ptr("repositories/myrepository/contentWrite"),
		// 					to.Ptr("repositories/myrepository/delete")},
		// 					CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
