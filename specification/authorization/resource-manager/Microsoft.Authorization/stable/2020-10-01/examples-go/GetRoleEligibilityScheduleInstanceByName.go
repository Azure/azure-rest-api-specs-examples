package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleInstanceByName.json
func ExampleRoleEligibilityScheduleInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleEligibilityScheduleInstancesClient().Get(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "21e4b59a-0499-4fe0-a3c3-43a3055b773a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleEligibilityScheduleInstance = armauthorization.RoleEligibilityScheduleInstance{
	// 	Name: to.Ptr("21e4b59a-0499-4fe0-a3c3-43a3055b773a"),
	// 	Type: to.Ptr("Microsoft.Authorization/RoleEligibilityScheduleInstances"),
	// 	ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilityScheduleInstances/21e4b59a-0499-4fe0-a3c3-43a3055b773a"),
	// 	Properties: &armauthorization.RoleEligibilityScheduleInstanceProperties{
	// 		Condition: to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
	// 		ConditionVersion: to.Ptr("1.0"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-10T00:32:36.860Z"); return t}()),
	// 		EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T00:31:41.477Z"); return t}()),
	// 		ExpandedProperties: &armauthorization.ExpandedProperties{
	// 			Principal: &armauthorization.ExpandedPropertiesPrincipal{
	// 				Type: to.Ptr("User"),
	// 				DisplayName: to.Ptr("User Account"),
	// 				Email: to.Ptr("user@my-tenant.com"),
	// 				ID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 			},
	// 			RoleDefinition: &armauthorization.ExpandedPropertiesRoleDefinition{
	// 				Type: to.Ptr("BuiltInRole"),
	// 				DisplayName: to.Ptr("Contributor"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
	// 			},
	// 			Scope: &armauthorization.ExpandedPropertiesScope{
	// 				Type: to.Ptr("subscription"),
	// 				DisplayName: to.Ptr("Pay-As-You-Go"),
	// 				ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 			},
	// 		},
	// 		MemberType: to.Ptr(armauthorization.MemberTypeDirect),
	// 		PrincipalID: to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c"),
	// 		RoleEligibilityScheduleID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/RoleEligibilitySchedules/b1477448-2cc6-4ceb-93b4-54a202a89413"),
	// 		Scope: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-10T00:32:36.860Z"); return t}()),
	// 		Status: to.Ptr(armauthorization.StatusProvisioned),
	// 	},
	// }
}
