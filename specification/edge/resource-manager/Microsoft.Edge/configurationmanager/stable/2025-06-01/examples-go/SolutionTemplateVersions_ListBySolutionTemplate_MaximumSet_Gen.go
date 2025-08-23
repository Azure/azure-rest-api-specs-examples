package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/SolutionTemplateVersions_ListBySolutionTemplate_MaximumSet_Gen.json
func ExampleSolutionTemplateVersionsClient_NewListBySolutionTemplatePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSolutionTemplateVersionsClient().NewListBySolutionTemplatePager("rgconfigurationmanager", "testname", nil)
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
		// page = armworkloadorchestration.SolutionTemplateVersionsClientListBySolutionTemplateResponse{
		// 	SolutionTemplateVersionListResult: armworkloadorchestration.SolutionTemplateVersionListResult{
		// 		Value: []*armworkloadorchestration.SolutionTemplateVersion{
		// 			{
		// 				Properties: &armworkloadorchestration.SolutionTemplateVersionProperties{
		// 					Configurations: to.Ptr("ofqcsavwmeuwmvtjnqpoybtjvkmrlh"),
		// 					Specification: map[string]any{
		// 					},
		// 					OrchestratorType: to.Ptr(armworkloadorchestration.OrchestratorTypeTO),
		// 					ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
		// 				},
		// 				ETag: to.Ptr("diqzzjtqycbtjzinckqwxowzdy"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("haebwdswymjzjfwbfipf"),
		// 				Type: to.Ptr("swo"),
		// 				SystemData: &armworkloadorchestration.SystemData{
		// 					CreatedBy: to.Ptr("nvjczgdguyvllp"),
		// 					CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
		// 					LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
