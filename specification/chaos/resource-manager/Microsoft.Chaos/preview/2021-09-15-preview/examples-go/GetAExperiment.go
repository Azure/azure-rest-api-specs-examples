package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/GetAExperiment.json
func ExampleExperimentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().Get(ctx, "exampleRG", "exampleExperiment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Experiment = armchaos.Experiment{
	// 	Name: to.Ptr("exampleExperiment"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Identity: &armchaos.ResourceIdentity{
	// 		Type: to.Ptr(armchaos.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
	// 		TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
	// 	},
	// 	Properties: &armchaos.ExperimentProperties{
	// 		Selectors: []*armchaos.Selector{
	// 			{
	// 				Type: to.Ptr(armchaos.SelectorTypeList),
	// 				ID: to.Ptr("selector1"),
	// 				Targets: []*armchaos.TargetReference{
	// 					{
	// 						Type: to.Ptr("ChaosTarget"),
	// 						ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"),
	// 				}},
	// 		}},
	// 		Steps: []*armchaos.Step{
	// 			{
	// 				Name: to.Ptr("step1"),
	// 				Branches: []*armchaos.Branch{
	// 					{
	// 						Name: to.Ptr("branch1"),
	// 						Actions: []armchaos.ActionClassification{
	// 							&armchaos.Action{
	// 								Name: to.Ptr("urn:csci:provider:providername:Shutdown/1.0"),
	// 								Type: to.Ptr("Continuous"),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// 	SystemData: &armchaos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User"),
	// 		LastModifiedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 	},
	// }
}
