package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3855ffb4be0cd4d227b130b67d874fa816736c04/specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/RedisEnterpriseAccessPolicyAssignmentGet.json
func ExampleAccessPolicyAssignmentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.AccessPolicyAssignment = armredisenterprise.AccessPolicyAssignment{
	// 	Name: to.Ptr("accessPolicyAssignmentName1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
	// 	ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 	Properties: &armredisenterprise.AccessPolicyAssignmentProperties{
	// 		AccessPolicyName: to.Ptr("default"),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		User: &armredisenterprise.AccessPolicyAssignmentPropertiesUser{
	// 			ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 		},
	// 	},
	// }
}
