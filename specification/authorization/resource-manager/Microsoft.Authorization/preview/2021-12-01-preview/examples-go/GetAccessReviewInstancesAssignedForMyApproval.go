package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewInstancesAssignedForMyApproval.json
func ExampleAccessReviewInstancesAssignedForMyApprovalClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccessReviewInstancesAssignedForMyApprovalClient().NewListPager("488a6d0e-0a63-4946-86e3-1f5bbc934661", &armauthorization.AccessReviewInstancesAssignedForMyApprovalClientListOptions{Filter: to.Ptr("assignedToMeToReview()")})
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
		// page.AccessReviewInstanceListResult = armauthorization.AccessReviewInstanceListResult{
		// 	Value: []*armauthorization.AccessReviewInstance{
		// 		{
		// 			Name: to.Ptr("4135f961-be78-4005-8101-c72a5af307a2"),
		// 			Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions/instances"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/488a6d0e-0a63-4946-86e3-1f5bbc934661/instances/4135f961-be78-4005-8101-c72a5af307a2"),
		// 			Properties: &armauthorization.AccessReviewInstanceProperties{
		// 				EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:17:30.513Z"); return t}()),
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
		// 				Status: to.Ptr(armauthorization.AccessReviewInstanceStatusApplied),
		// 			},
		// 	}},
		// }
	}
}
