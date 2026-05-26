package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/ResiliencyReviewsGet.json
func ExampleResiliencyReviewsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResiliencyReviewsClient().Get(ctx, "11111111-1111-2222-3333-444444444444", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armadvisor.ResiliencyReviewsClientGetResponse{
	// 	ResiliencyReview: armadvisor.ResiliencyReview{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Advisor/ResiliencyReview/11111111-1111-2222-3333-444444444444"),
	// 		Name: to.Ptr("11111111-1111-2222-3333-444444444444"),
	// 		Type: to.Ptr("Microsoft.Advisor/ResiliencyReview"),
	// 		SystemData: &armadvisor.SystemData{
	// 			CreatedBy: to.Ptr("user-identity"),
	// 			CreatedByType: to.Ptr(armadvisor.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-30T09:41:00Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user-identity"),
	// 			LastModifiedByType: to.Ptr(armadvisor.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-30T09:41:00Z"); return t}()),
	// 		},
	// 		Properties: &armadvisor.ResiliencyReviewProperties{
	// 			ReviewName: to.Ptr("review name xyz"),
	// 			WorkloadName: to.Ptr("workload name xyz"),
	// 			ReviewStatus: to.Ptr(armadvisor.ReviewStatusTriaged),
	// 			RecommendationsCount: to.Ptr[int32](4),
	// 			PublishedAt: to.Ptr("2023-06-30T09:41:00Z"),
	// 			UpdatedAt: to.Ptr("2023-06-30T09:41:00Z"),
	// 		},
	// 	},
	// }
}
