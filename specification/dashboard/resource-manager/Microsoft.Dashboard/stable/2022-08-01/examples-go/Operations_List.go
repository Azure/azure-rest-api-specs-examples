package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armdashboard.OperationListResult{
		// 	Value: []*armdashboard.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Dashboard/grafana/write"),
		// 			ActionType: to.Ptr(armdashboard.ActionTypeInternal),
		// 			Display: &armdashboard.OperationDisplay{
		// 				Description: to.Ptr("Write grafana workspace resource"),
		// 				Operation: to.Ptr("write"),
		// 				Provider: to.Ptr("Microsoft.Dashboard"),
		// 				Resource: to.Ptr("grafana"),
		// 			},
		// 			Origin: to.Ptr(armdashboard.OriginUser),
		// 	}},
		// }
	}
}
