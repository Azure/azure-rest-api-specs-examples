package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2026-03-01-preview/Projects_TierLimitReached_MaximumSet_Gen.json
func ExampleProjectsClient_TierLimitReached() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("1E4BD993-6890-4E69-8966-81482D7502EF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().TierLimitReached(ctx, "rgopenapi", "myOrganization", "myProject", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongodbatlas.ProjectsClientTierLimitReachedResponse{
	// 	TierLimitReachedResponse: armmongodbatlas.TierLimitReachedResponse{
	// 		Limits: []*armmongodbatlas.ProjectLimitStatus{
	// 			{
	// 				Type: to.Ptr(armmongodbatlas.ClusterTierFREE),
	// 				Maximum: to.Ptr[int32](29),
	// 				Current: to.Ptr[int32](27),
	// 				IsReached: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
