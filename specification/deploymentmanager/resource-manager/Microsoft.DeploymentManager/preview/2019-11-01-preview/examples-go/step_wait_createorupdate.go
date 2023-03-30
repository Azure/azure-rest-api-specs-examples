package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_wait_createorupdate.json
func ExampleStepsClient_CreateOrUpdate_createWaitStep() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStepsClient().CreateOrUpdate(ctx, "myResourceGroup", "waitStep", &armdeploymentmanager.StepsClientCreateOrUpdateOptions{StepInfo: &armdeploymentmanager.StepResource{
		Location: to.Ptr("centralus"),
		Tags:     map[string]*string{},
		Properties: &armdeploymentmanager.WaitStepProperties{
			StepType: to.Ptr(armdeploymentmanager.StepTypeWait),
			Attributes: &armdeploymentmanager.WaitStepAttributes{
				Duration: to.Ptr("PT20M"),
			},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
