package armauthorization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PutRoleEligibilityScheduleRequest.json
func ExampleRoleEligibilityScheduleRequestsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRoleEligibilityScheduleRequestsClient().Create(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			Condition:        to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
			ConditionVersion: to.Ptr("1.0"),
			PrincipalID:      to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
			RequestType:      to.Ptr(armauthorization.RequestTypeAdminAssign),
			RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
			ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: to.Ptr("P365D"),
				},
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.910Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
