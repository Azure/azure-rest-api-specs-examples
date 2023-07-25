package armauthorization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-04-01-preview/examples/PutRoleAssignmentScheduleRequest.json
func ExampleRoleAssignmentScheduleRequestsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRoleAssignmentScheduleRequestsClient().Create(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "fea7a502-9a96-4806-a26f-eee560e52045", armauthorization.RoleAssignmentScheduleRequest{
		Properties: &armauthorization.RoleAssignmentScheduleRequestProperties{
			Condition:                       to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
			ConditionVersion:                to.Ptr("1.0"),
			LinkedRoleEligibilityScheduleID: to.Ptr("b1477448-2cc6-4ceb-93b4-54a202a89413"),
			PrincipalID:                     to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
			RequestType:                     to.Ptr(armauthorization.RequestTypeSelfActivate),
			RoleDefinitionID:                to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
			ScheduleInfo: &armauthorization.RoleAssignmentScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: to.Ptr("PT8H"),
				},
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:35:27.91Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
