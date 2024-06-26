package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewHistoryDefinitions.json
func ExampleScopeAccessReviewHistoryDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScopeAccessReviewHistoryDefinitionsClient().NewListPager("subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a", &armauthorization.ScopeAccessReviewHistoryDefinitionsClientListOptions{Filter: nil})
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
		// page.AccessReviewHistoryDefinitionListResult = armauthorization.AccessReviewHistoryDefinitionListResult{
		// 	Value: []*armauthorization.AccessReviewHistoryDefinition{
		// 		{
		// 			Name: to.Ptr("44724910-d7a5-4c29-b28f-db73e717165a"),
		// 			Type: to.Ptr("Microsoft.Authorization/accessReviewHistoryDefinition"),
		// 			ID: to.Ptr("/subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/44724910-d7a5-4c29-b28f-db73e717165a"),
		// 			Properties: &armauthorization.AccessReviewHistoryDefinitionProperties{
		// 				CreatedBy: &armauthorization.AccessReviewActorIdentity{
		// 					PrincipalID: to.Ptr("673ad0d8-7b0e-4201-aaeb-74cdcbf22af9"),
		// 					PrincipalName: to.Ptr("levi"),
		// 					UserPrincipalName: to.Ptr("levi"),
		// 				},
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-23T00:27:33.690Z"); return t}()),
		// 				Decisions: []*armauthorization.AccessReviewResult{
		// 					to.Ptr(armauthorization.AccessReviewResultApprove),
		// 					to.Ptr(armauthorization.AccessReviewResultDeny),
		// 					to.Ptr(armauthorization.AccessReviewResultNotReviewed),
		// 					to.Ptr(armauthorization.AccessReviewResultDontKnow),
		// 					to.Ptr(armauthorization.AccessReviewResultNotNotified)},
		// 					DisplayName: to.Ptr("Hello world name"),
		// 					ReviewHistoryPeriodEndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-01T08:00:00.000Z"); return t}()),
		// 					ReviewHistoryPeriodStartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T07:00:00.000Z"); return t}()),
		// 					Scopes: []*armauthorization.AccessReviewScope{
		// 						{
		// 							ExpandNestedMemberships: to.Ptr(true),
		// 							PrincipalType: to.Ptr(armauthorization.AccessReviewScopePrincipalTypeUser),
		// 							ResourceID: to.Ptr("/subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a"),
		// 							RoleDefinitionID: to.Ptr("/subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a/providers/Microsoft.Authorization/roleDefinitions"),
		// 						},
		// 						{
		// 							ExpandNestedMemberships: to.Ptr(false),
		// 							PrincipalType: to.Ptr(armauthorization.AccessReviewScopePrincipalTypeUser),
		// 							ResourceID: to.Ptr("/subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a"),
		// 							RoleDefinitionID: to.Ptr("/subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a/providers/Microsoft.Authorization/roleDefinitions/1562cf42-00b9-457c-86ef-5702d4132904"),
		// 					}},
		// 					Settings: &armauthorization.AccessReviewHistoryScheduleSettings{
		// 						Range: &armauthorization.AccessReviewRecurrenceRange{
		// 							Type: to.Ptr(armauthorization.AccessReviewRecurrenceRangeTypeEndDate),
		// 							EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-01T08:00:00.000Z"); return t}()),
		// 							NumberOfOccurrences: to.Ptr[int32](1),
		// 							StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-01T08:00:00.000Z"); return t}()),
		// 						},
		// 					},
		// 					Status: to.Ptr(armauthorization.AccessReviewHistoryDefinitionStatusDone),
		// 				},
		// 		}},
		// 	}
	}
}
