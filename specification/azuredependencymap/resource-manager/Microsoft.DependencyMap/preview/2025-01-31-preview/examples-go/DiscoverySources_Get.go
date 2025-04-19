package armdependencymap_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dependencymap/armdependencymap"
)

// Generated from example definition: 2025-01-31-preview/DiscoverySources_Get.json
func ExampleDiscoverySourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdependencymap.NewClientFactory("D6E58BDB-45F1-41EC-A884-1FC945058848", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiscoverySourcesClient().Get(ctx, "rgdependencyMap", "mapsTest1", "sourceTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdependencymap.DiscoverySourcesClientGetResponse{
	// 	DiscoverySourceResource: &armdependencymap.DiscoverySourceResource{
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("y"),
	// 		ID: to.Ptr("hgq"),
	// 		Name: to.Ptr("opthbwbxhckrctnfnpio"),
	// 		Type: to.Ptr("blxnetzeqhqslnbokqnsaib"),
	// 		SystemData: &armdependencymap.SystemData{
	// 			CreatedBy: to.Ptr("dgycvdwmcmlllzqi"),
	// 			CreatedByType: to.Ptr(armdependencymap.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-29T07:35:13.318Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("wfntygzyylkddshm"),
	// 			LastModifiedByType: to.Ptr(armdependencymap.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-29T07:35:13.318Z"); return t}()),
	// 		},
	// 	},
	// }
}
