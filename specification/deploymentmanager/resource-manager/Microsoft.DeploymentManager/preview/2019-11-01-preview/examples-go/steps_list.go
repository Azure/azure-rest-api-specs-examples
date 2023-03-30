package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/steps_list.json
func ExampleStepsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStepsClient().List(ctx, "myResourceGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StepResourceArray = []*armdeploymentmanager.StepResource{
	// 	{
	// 		Name: to.Ptr("waitStepfrontEnd"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/steps"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.WaitStepProperties{
	// 			StepType: to.Ptr(armdeploymentmanager.StepTypeWait),
	// 			Attributes: &armdeploymentmanager.WaitStepAttributes{
	// 				Duration: to.Ptr("PT10M"),
	// 			},
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("waitStepBackEnd"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/steps"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.WaitStepProperties{
	// 			StepType: to.Ptr(armdeploymentmanager.StepTypeWait),
	// 			Attributes: &armdeploymentmanager.WaitStepAttributes{
	// 				Duration: to.Ptr("PT30M"),
	// 			},
	// 		},
	// }}
}
