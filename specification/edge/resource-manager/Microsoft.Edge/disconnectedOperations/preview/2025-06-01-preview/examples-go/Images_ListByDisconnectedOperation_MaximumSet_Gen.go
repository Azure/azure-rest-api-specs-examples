package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/Images_ListByDisconnectedOperation_MaximumSet_Gen.json
func ExampleImagesClient_NewListByDisconnectedOperationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("1F6CACA0-5FFA-47AD-94FD-42368F71E49E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImagesClient().NewListByDisconnectedOperationPager("rgdisconnectedoperations", "w_-EG-3-euL7K3-E", &armdisconnectedoperations.ImagesClientListByDisconnectedOperationOptions{
		Filter: to.Ptr("toynendoobwkrcwmfdfup"),
		Top:    to.Ptr[int32](20),
		Skip:   to.Ptr[int32](3)})
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
		// page = armdisconnectedoperations.ImagesClientListByDisconnectedOperationResponse{
		// 	ImageListResult: armdisconnectedoperations.ImageListResult{
		// 		Value: []*armdisconnectedoperations.Image{
		// 			{
		// 				Properties: &armdisconnectedoperations.ImageProperties{
		// 					ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
		// 					ReleaseVersion: to.Ptr("2.0.0"),
		// 					ReleaseDisplayName: to.Ptr("release 1"),
		// 					ReleaseNotes: to.Ptr("notes"),
		// 					ReleaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.DateOnly, "2025-05-14"); return t}()),
		// 					ReleaseType: to.Ptr(armdisconnectedoperations.ReleaseTypeInstall),
		// 					CompatibleVersions: []*string{
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/1F6CACA0-5FFA-47AD-94FD-42368F71E49E/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
		// 				Name: to.Ptr("demo-resource"),
		// 				Type: to.Ptr("Microsoft.Edge/disconnectedOperations/images"),
		// 				SystemData: &armdisconnectedoperations.SystemData{
		// 					CreatedBy: to.Ptr("bwpmoygkcv"),
		// 					CreatedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-14T19:03:52.617Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("tottbcqgtvsjebkejjlhxibnaifijd"),
		// 					LastModifiedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-14T19:03:52.617Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aan"),
		// 	},
		// }
	}
}
