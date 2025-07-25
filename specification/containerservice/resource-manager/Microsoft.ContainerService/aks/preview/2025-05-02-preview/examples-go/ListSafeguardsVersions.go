package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97f789aeb52adfc1e20c386005839f5276874d7d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-05-02-preview/examples/ListSafeguardsVersions.json
func ExampleManagedClustersClient_NewListSafeguardsVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListSafeguardsVersionsPager("location1", nil)
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
		// page.SafeguardsAvailableVersionsList = armcontainerservice.SafeguardsAvailableVersionsList{
		// 	Value: []*armcontainerservice.SafeguardsAvailableVersion{
		// 		{
		// 			Name: to.Ptr("v1.0.0"),
		// 			Type: to.Ptr("Microsoft.ContainerService/locations/safeguardsVersions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ContainerService/locations/location1/safeguardsVersions/v1.0.0"),
		// 			Properties: &armcontainerservice.SafeguardsAvailableVersionsProperties{
		// 				IsDefaultVersion: to.Ptr(true),
		// 				Support: to.Ptr(armcontainerservice.SafeguardsSupportPreview),
		// 			},
		// 	}},
		// }
	}
}
