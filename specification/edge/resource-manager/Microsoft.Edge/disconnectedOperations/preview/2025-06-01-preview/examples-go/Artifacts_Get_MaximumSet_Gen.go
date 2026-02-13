package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/Artifacts_Get_MaximumSet_Gen.json
func ExampleArtifactsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("301DCB09-82EC-4777-A56C-6FFF26BCC814", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactsClient().Get(ctx, "rgdisconnectedoperations", "J_3-_S--_-UM_-_7w11", "PMY-", "-8Y-Us1BNNG6-H5w6-2--RP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdisconnectedoperations.ArtifactsClientGetResponse{
	// 	Artifact: &armdisconnectedoperations.Artifact{
	// 		Properties: &armdisconnectedoperations.ArtifactProperties{
	// 			ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
	// 			ArtifactOrder: to.Ptr[int32](1),
	// 			Title: to.Ptr("artifact pat 1"),
	// 			Description: to.Ptr("the first part of the image"),
	// 			Size: to.Ptr[int64](22),
	// 		},
	// 		ID: to.Ptr("/subscriptions/1F6CACA0-5FFA-47AD-94FD-42368F71E49E/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
	// 		Name: to.Ptr("demo-resource"),
	// 		Type: to.Ptr("Microsoft.Edge/disconnectedOperations/images/artifacts"),
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
