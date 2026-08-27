package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2026-06-01-preview/RedisEnterpriseAccessPolicyAssignmentGetWithFailedProvisioningState.json
func ExampleAccessPolicyAssignmentClient_Get_redisEnterpriseAccessPolicyAssignmentGetWithFailedProvisioningState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessPolicyAssignmentClient().Get(ctx, "rg1", "cache1", "default", "accessPolicyAssignmentName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredisenterprise.AccessPolicyAssignmentClientGetResponse{
	// 	AccessPolicyAssignment: armredisenterprise.AccessPolicyAssignment{
	// 		Name: to.Ptr("accessPolicyAssignmentName1"),
	// 		Type: to.Ptr("Microsoft.Cache/redisEnterprise/databases/accessPolicyAssignments"),
	// 		ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 		Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
	// 			AccessPolicyName: to.Ptr("default"),
	// 			AccessString: to.Ptr("+@invalid-syntax"),
	// 			ProvisioningError: &armredisenterprise.AccessPolicyAssignmentProvisioningError{
	// 				Code: to.Ptr("InvalidAccessString"),
	// 				Message: to.Ptr("ERR Error in ACL SETUSER modifier '+@invalid-syntax': Adding a subcommand of a command already fully added is not allowed. A full command can only be added to a rule"),
	// 				Target: to.Ptr("properties.accessString"),
	// 			},
	// 			ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateFailed),
	// 			User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
	// 				ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 			},
	// 		},
	// 	},
	// }
}
