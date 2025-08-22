package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/SolutionTemplates_Get_MaximumSet_Gen.json
func ExampleSolutionTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionTemplatesClient().Get(ctx, "rgconfigurationmanager", "testname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadorchestration.SolutionTemplatesClientGetResponse{
	// 	SolutionTemplate: &armworkloadorchestration.SolutionTemplate{
	// 		Properties: &armworkloadorchestration.SolutionTemplateProperties{
	// 			Description: to.Ptr("psrftehgzngcdlccivhjmwsmiz"),
	// 			Capabilities: []*string{
	// 				to.Ptr("dfoyxbbknrhvlunhmuyyt"),
	// 			},
	// 			LatestVersion: to.Ptr("vk"),
	// 			State: to.Ptr(armworkloadorchestration.ResourceStateActive),
	// 			EnableExternalValidation: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		ETag: to.Ptr("zotshwhsdrfk"),
	// 		Tags: map[string]*string{
	// 			"key5091": to.Ptr("dov"),
	// 		},
	// 		Location: to.Ptr("zheaaqvadewftnctxzpinrgeproqs"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("hhmslfrwpuvrjr"),
	// 		Type: to.Ptr("qxqlylfzyiykcdsbwuanmgrqwlfgjo"),
	// 		SystemData: &armworkloadorchestration.SystemData{
	// 			CreatedBy: to.Ptr("nvjczgdguyvllp"),
	// 			CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
	// 			LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 		},
	// 	},
	// }
}
