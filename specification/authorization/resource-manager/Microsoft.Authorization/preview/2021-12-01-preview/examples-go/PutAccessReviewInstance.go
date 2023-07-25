package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/PutAccessReviewInstance.json
func ExampleScopeAccessReviewInstancesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopeAccessReviewInstancesClient().Create(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "4135f961-be78-4005-8101-c72a5af307a2", armauthorization.AccessReviewInstanceProperties{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessReviewInstance = armauthorization.AccessReviewInstance{
	// 	Name: to.Ptr("4135f961-be78-4005-8101-c72a5af307a2"),
	// 	Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions/instances"),
	// 	ID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/instances/4135f961-be78-4005-8101-c72a5af307a2"),
	// 	Properties: &armauthorization.AccessReviewInstanceProperties{
	// 		BackupReviewers: []*armauthorization.AccessReviewReviewer{
	// 			{
	// 				PrincipalID: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 				PrincipalType: to.Ptr(armauthorization.AccessReviewReviewerTypeUser),
	// 		}},
	// 		EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:17:30.513Z"); return t}()),
	// 		Reviewers: []*armauthorization.AccessReviewReviewer{
	// 			{
	// 				PrincipalID: to.Ptr("89ff94f6-029c-429c-9037-6cf4fe2f4a13"),
	// 				PrincipalType: to.Ptr(armauthorization.AccessReviewReviewerType("group")),
	// 		}},
	// 		ReviewersType: to.Ptr(armauthorization.AccessReviewInstanceReviewersTypeAssigned),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
	// 		Status: to.Ptr(armauthorization.AccessReviewInstanceStatusInProgress),
	// 	},
	// }
}
