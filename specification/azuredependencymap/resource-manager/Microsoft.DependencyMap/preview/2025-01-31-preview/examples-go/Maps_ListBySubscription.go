package armdependencymap_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dependencymap/armdependencymap"
)

// Generated from example definition: 2025-01-31-preview/Maps_ListBySubscription.json
func ExampleMapsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdependencymap.NewClientFactory("D6E58BDB-45F1-41EC-A884-1FC945058848", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMapsClient().NewListBySubscriptionPager(nil)
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
		// page = armdependencymap.MapsClientListBySubscriptionResponse{
		// 	MapsResourceListResult: armdependencymap.MapsResourceListResult{
		// 		Value: []*armdependencymap.MapsResource{
		// 			{
		// 				Properties: &armdependencymap.MapsResourceProperties{
		// 					ProvisioningState: to.Ptr(armdependencymap.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("wjtzelgfcmswfeflfltwxqveo"),
		// 				ID: to.Ptr("oiaolujazzjebav"),
		// 				Name: to.Ptr("pwkdwgji"),
		// 				Type: to.Ptr("cxtymeelrjudiduuwi"),
		// 				SystemData: &armdependencymap.SystemData{
		// 					CreatedBy: to.Ptr("dgycvdwmcmlllzqi"),
		// 					CreatedByType: to.Ptr(armdependencymap.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-29T07:35:13.318Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("wfntygzyylkddshm"),
		// 					LastModifiedByType: to.Ptr(armdependencymap.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-29T07:35:13.318Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
