package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceRunners_Get.json
func ExampleServiceRunnersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceRunnersClient().Get(ctx, "resourceGroupName", "{devtestlabName}", "{servicerunnerName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceRunner = armdevtestlabs.ServiceRunner{
	// 	Name: to.Ptr("{serviceRunnerName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/serviceRunners"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/serviceRunners/{serviceRunnerName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Identity: &armdevtestlabs.IdentityProperties{
	// 		Type: to.Ptr(armdevtestlabs.ManagedIdentityType("{identityType}")),
	// 		ClientSecretURL: to.Ptr("{identityClientSecretUrl}"),
	// 		PrincipalID: to.Ptr("{identityPrincipalId}"),
	// 		TenantID: to.Ptr("{identityTenantId}"),
	// 	},
	// }
}
