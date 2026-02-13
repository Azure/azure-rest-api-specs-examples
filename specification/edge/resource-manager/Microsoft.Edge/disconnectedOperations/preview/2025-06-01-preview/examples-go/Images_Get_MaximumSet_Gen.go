package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/Images_Get_MaximumSet_Gen.json
func ExampleImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("301DCB09-82EC-4777-A56C-6FFF26BCC814", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImagesClient().Get(ctx, "rgdisconnectedoperations", "bT62l-KS7g1-uh", "2P6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdisconnectedoperations.ImagesClientGetResponse{
	// 	Image: &armdisconnectedoperations.Image{
	// 		Properties: &armdisconnectedoperations.ImageProperties{
	// 			ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
	// 			ReleaseVersion: to.Ptr("2.0.0"),
	// 			ReleaseDisplayName: to.Ptr("release 1"),
	// 			ReleaseNotes: to.Ptr("notes"),
	// 			ReleaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.DateOnly, "2025-05-14"); return t}()),
	// 			ReleaseType: to.Ptr(armdisconnectedoperations.ReleaseTypeInstall),
	// 			CompatibleVersions: []*string{
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/1F6CACA0-5FFA-47AD-94FD-42368F71E49E/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
	// 		Name: to.Ptr("demo-resource"),
	// 		Type: to.Ptr("Microsoft.Edge/disconnectedOperations/images"),
	// 		SystemData: &armdisconnectedoperations.SystemData{
	// 			CreatedBy: to.Ptr("bwpmoygkcv"),
	// 			CreatedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-14T19:03:52.617Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("tottbcqgtvsjebkejjlhxibnaifijd"),
	// 			LastModifiedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-14T19:03:52.617Z"); return t}()),
	// 		},
	// 	},
	// }
}
