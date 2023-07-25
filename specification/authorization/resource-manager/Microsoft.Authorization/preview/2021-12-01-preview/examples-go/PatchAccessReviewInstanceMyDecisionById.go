package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/PatchAccessReviewInstanceMyDecisionById.json
func ExampleAccessReviewInstanceMyDecisionsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessReviewInstanceMyDecisionsClient().Patch(ctx, "488a6d0e-0a63-4946-86e3-1f5bbc934661", "4135f961-be78-4005-8101-c72a5af307a2", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", armauthorization.AccessReviewDecisionProperties{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessReviewDecision = armauthorization.AccessReviewDecision{
	// 	Name: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 	Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions/instances/decisions"),
	// 	ID: to.Ptr("/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/488a6d0e-0a63-4946-86e3-1f5bbc934661/instances/4135f961-be78-4005-8101-c72a5af307a2/decisions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 	Properties: &armauthorization.AccessReviewDecisionProperties{
	// 		AppliedBy: &armauthorization.AccessReviewActorIdentity{
	// 			PrincipalID: to.Ptr("36777fc8-4ec2-49ea-a56c-cec0bd47d83a"),
	// 			PrincipalName: to.Ptr("Amit Ghosh"),
	// 			PrincipalType: to.Ptr(armauthorization.AccessReviewActorIdentityTypeUser),
	// 			UserPrincipalName: to.Ptr("amitgho@microsoft.com"),
	// 		},
	// 		AppliedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-03T21:02:30.667Z"); return t}()),
	// 		ApplyResult: to.Ptr(armauthorization.AccessReviewApplyResult("Success")),
	// 		Decision: to.Ptr(armauthorization.AccessReviewResultDeny),
	// 		Justification: to.Ptr("This person has left this team"),
	// 		Principal: &armauthorization.AccessReviewDecisionUserIdentity{
	// 			Type: to.Ptr(armauthorization.DecisionTargetTypeUser),
	// 			DisplayName: to.Ptr("Shubham Gupta"),
	// 			ID: to.Ptr("a6c7aecb-cbfd-4763-87ef-e91b4bd509d9"),
	// 			UserPrincipalName: to.Ptr("shugup@microsoft.com"),
	// 		},
	// 		Recommendation: to.Ptr(armauthorization.AccessRecommendationTypeDeny),
	// 		Resource: &armauthorization.AccessReviewDecisionResource{
	// 			Type: to.Ptr(armauthorization.DecisionResourceTypeAzureRole),
	// 			DisplayName: to.Ptr("Owner"),
	// 			ID: to.Ptr("a6c7aecb-cbfd-4763-87ef-e91b4bd509d9"),
	// 		},
	// 		ReviewedBy: &armauthorization.AccessReviewActorIdentity{
	// 			PrincipalID: to.Ptr("a6c7aecb-cbfd-4763-87ef-e91b4bd509d9"),
	// 			PrincipalName: to.Ptr("Shubham Gupta"),
	// 			PrincipalType: to.Ptr(armauthorization.AccessReviewActorIdentityTypeUser),
	// 			UserPrincipalName: to.Ptr("shugup@microsoft.com"),
	// 		},
	// 		ReviewedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
	// 	},
	// }
}
