package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewDefaultSettings.json
func ExampleScopeAccessReviewDefaultSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopeAccessReviewDefaultSettingsClient().Get(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessReviewDefaultSettings = armauthorization.AccessReviewDefaultSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Authorization/accessReviewSettings"),
	// 	ID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/accessReviewSettings/default"),
	// 	Properties: &armauthorization.AccessReviewScheduleSettings{
	// 		AutoApplyDecisionsEnabled: to.Ptr(true),
	// 		DefaultDecision: to.Ptr(armauthorization.DefaultDecisionTypeApprove),
	// 		DefaultDecisionEnabled: to.Ptr(true),
	// 		InstanceDurationInDays: to.Ptr[int32](30),
	// 		JustificationRequiredOnApproval: to.Ptr(true),
	// 		MailNotificationsEnabled: to.Ptr(true),
	// 		RecommendationsEnabled: to.Ptr(true),
	// 		Recurrence: &armauthorization.AccessReviewRecurrenceSettings{
	// 			Range: &armauthorization.AccessReviewRecurrenceRange{
	// 				Type: to.Ptr(armauthorization.AccessReviewRecurrenceRangeTypeEndDate),
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:17:30.513Z"); return t}()),
	// 				NumberOfOccurrences: to.Ptr[int32](1),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
	// 			},
	// 		},
	// 		ReminderNotificationsEnabled: to.Ptr(true),
	// 	},
	// }
}
