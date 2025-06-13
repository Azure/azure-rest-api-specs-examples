package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2025-01-01/Targets_List.json
func ExampleTargetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTargetsClient().NewListPager("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", nil)
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
		// page = armchaos.TargetsClientListResponse{
		// 	TargetListResult: armchaos.TargetListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets?continuationToken=&api-version=2024-11-01-preview"),
		// 		Value: []*armchaos.Target{
		// 			{
		// 				Name: to.Ptr("Microsoft-Agent"),
		// 				Type: to.Ptr("Microsoft.Chaos/targets"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-Agent"),
		// 				Location: to.Ptr("centraluseuap"),
		// 				Properties: map[string]any{
		// 					"agentProfileId": "ac4e8251-fdc9-4277-8e87-dc57fe5794cf",
		// 					"identities": []any{
		// 						map[string]any{
		// 							"type": "CertificateSubjectIssuer",
		// 							"subject": "CN=example.subject",
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
