package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/Artifact_ListByParent_MaximumSet_Gen.json
func ExampleArtifactsClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("301DCB09-82EC-4777-A56C-6FFF26BCC814", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArtifactsClient().NewListByParentPager("rgdisconnectedoperations", "XOn_Y-7_M-46E-Y", "2v5Q3mNihPV88C882LnbQO8", nil)
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
		// page = armdisconnectedoperations.ArtifactsClientListByParentResponse{
		// 	ArtifactListResult: armdisconnectedoperations.ArtifactListResult{
		// 		Value: []*armdisconnectedoperations.Artifact{
		// 			{
		// 				Properties: &armdisconnectedoperations.ArtifactProperties{
		// 					ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
		// 					ArtifactOrder: to.Ptr[int32](1),
		// 					Title: to.Ptr("artifact pat 1"),
		// 					Description: to.Ptr("the first part of the image"),
		// 					Size: to.Ptr[int64](22),
		// 				},
		// 				ID: to.Ptr("/subscriptions/1F6CACA0-5FFA-47AD-94FD-42368F71E49E/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
		// 				Name: to.Ptr("demo-resource"),
		// 				Type: to.Ptr("Microsoft.Edge/disconnectedOperations/images/artifacts"),
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
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
