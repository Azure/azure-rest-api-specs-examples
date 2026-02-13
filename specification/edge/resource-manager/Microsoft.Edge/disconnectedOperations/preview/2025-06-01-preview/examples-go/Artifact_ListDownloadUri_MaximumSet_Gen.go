package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/Artifact_ListDownloadUri_MaximumSet_Gen.json
func ExampleArtifactsClient_ListDownloadURI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("301DCB09-82EC-4777-A56C-6FFF26BCC814", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactsClient().ListDownloadURI(ctx, "rgdisconnectedoperations", "L4z_-S", "B-Ra--W0", "artifact1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdisconnectedoperations.ArtifactsClientListDownloadURIResponse{
	// 	ArtifactDownloadResult: &armdisconnectedoperations.ArtifactDownloadResult{
	// 		ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
	// 		ArtifactOrder: to.Ptr[int32](1),
	// 		Title: to.Ptr("artifact pat 1"),
	// 		Description: to.Ptr("the first part of the image"),
	// 		Size: to.Ptr[int64](29000),
	// 		DownloadLink: to.Ptr("https://microsoft.com/a"),
	// 		LinkExpiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-14T19:03:53.740Z"); return t}()),
	// 	},
	// }
}
