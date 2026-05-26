package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/TriageRecommendationsGet.json
func ExampleTriageRecommendationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriageRecommendationsClient().Get(ctx, "11111111-1111-2222-3333-444444444444", "22222222-1111-2222-3333-444444444444", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armadvisor.TriageRecommendationsClientGetResponse{
	// 	TriageRecommendation: armadvisor.TriageRecommendation{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Advisor/ResiliencyReview/11111111-1111-2222-3333-444444444444/providers/Microsoft.Advisor/TriageRecommendations/22222222-1111-2222-3333-444444444444"),
	// 		Name: to.Ptr("22222222-1111-2222-3333-444444444444"),
	// 		Type: to.Ptr("Microsoft.Advisor/TriageRecommendations"),
	// 		SystemData: &armadvisor.SystemData{
	// 			CreatedBy: to.Ptr("user-identity"),
	// 			CreatedByType: to.Ptr(armadvisor.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-30T09:41:00Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user-identity"),
	// 			LastModifiedByType: to.Ptr(armadvisor.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-30T09:41:00Z"); return t}()),
	// 		},
	// 		Properties: &armadvisor.TriageRecommendationProperties{
	// 			ReviewID: to.Ptr("11111111-1111-2222-3333-444444444444"),
	// 			Title: to.Ptr("Explore the scale-limits of each Azure resource"),
	// 			Priority: to.Ptr(armadvisor.PriorityNameHigh),
	// 			RecommendationStatus: to.Ptr(armadvisor.RecommendationStatusNamePending),
	// 			UpdatedAt: to.Ptr("2023-06-30T10:51:00Z"),
	// 			RejectReason: to.Ptr(""),
	// 			AppliesToSubscriptions: []*string{
	// 				to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				to.Ptr("00000000-1111-2222-3333-555555555555"),
	// 				to.Ptr("00000000-1111-2222-3333-666666666666"),
	// 			},
	// 			Description: to.Ptr("Azure resources have scale limits. If you are planning to deploy a large number of resources, you should be aware of these limits. This recommendation provides links to the scale limits for each Azure resource."),
	// 			PotentialBenefits: to.Ptr("Avoid deployment failures due to scale limits."),
	// 			Notes: to.Ptr("This recommendation is based on the scale limits for each Azure resource. For more information, see <a href=\"https://docs.microsoft.com/azure/azure-resource-manager/management/azure-subscription-service-limits\">Azure subscription and service limits, quotas, and constraints</a>."),
	// 		},
	// 	},
	// }
}
