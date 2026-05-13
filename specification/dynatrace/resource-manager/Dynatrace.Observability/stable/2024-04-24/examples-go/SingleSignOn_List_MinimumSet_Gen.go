package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: 2024-04-24/SingleSignOn_List_MinimumSet_Gen.json
func ExampleSingleSignOnClient_NewListPager_singleSignOnListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSingleSignOnClient().NewListPager("myResourceGroup", "myMonitor", nil)
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
		// page = armdynatrace.SingleSignOnClientListResponse{
		// 	SingleSignOnResourceListResult: armdynatrace.SingleSignOnResourceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Dynatrace.Observability/monitors?api-version=2024-04-24&$skiptoken=abc123"),
		// 		Value: []*armdynatrace.SingleSignOnResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Dynatrace.Observability/monitors/myMonitor/singleSignOnConfigurations/default"),
		// 				Properties: &armdynatrace.SingleSignOnProperties{
		// 					AADDomains: []*string{
		// 						to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com"),
		// 					},
		// 					SingleSignOnURL: to.Ptr("https://www.dynatrace.io/IAmSomeHash"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
