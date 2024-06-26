package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751f321654db00858e70649291af5c81e94611e/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_List_MaximumSet_Gen.json
func ExampleSingleSignOnClient_NewListPager_singleSignOnListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.SingleSignOnResourceListResult = armdynatrace.SingleSignOnResourceListResult{
		// 	Value: []*armdynatrace.SingleSignOnResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Dynatrace.Observability/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Dynatrace.Observability/monitors/myMonitor/singleSignOnConfigurations/default"),
		// 			Properties: &armdynatrace.SingleSignOnProperties{
		// 				AADDomains: []*string{
		// 					to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
		// 					EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateSucceeded),
		// 					SingleSignOnState: to.Ptr(armdynatrace.SingleSignOnStatesEnable),
		// 					SingleSignOnURL: to.Ptr("https://www.dynatrace.io/IAmSomeHash"),
		// 				},
		// 				SystemData: &armdynatrace.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
		// 					CreatedBy: to.Ptr("alice@microsoft.com"),
		// 					CreatedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("alice@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}
