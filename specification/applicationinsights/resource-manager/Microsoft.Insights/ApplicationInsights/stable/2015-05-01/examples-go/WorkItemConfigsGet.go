package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: 2015-05-01/WorkItemConfigsGet.json
func ExampleWorkItemConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkItemConfigurationsClient().NewListPager("my-resource-group", "my-component", nil)
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
		// page = armapplicationinsights.WorkItemConfigurationsClientListResponse{
		// 	WorkItemConfigurationsListResult: armapplicationinsights.WorkItemConfigurationsListResult{
		// 		Value: []*armapplicationinsights.WorkItemConfiguration{
		// 			{
		// 				ConfigDisplayName: to.Ptr("Visual Studio Team Services"),
		// 				ConfigProperties: to.Ptr("{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"b370eeb2-5dff-4838-8253-ec3014b2c706\",\"ConfigurationType\":\"Standard\",\"WorkItemType\":\"Bug\",\"AreaPath\":\"todeletefirst\",\"AssignedTo\":\"\"}"),
		// 				ConnectorID: to.Ptr("d334e2a4-6733-488e-8645-a9fdc1694f41"),
		// 				ID: to.Ptr("Visual Studio Team Services"),
		// 				IsDefault: to.Ptr(true),
		// 			},
		// 			{
		// 				ConfigDisplayName: to.Ptr("GitHub"),
		// 				ConfigProperties: to.Ptr("{\"GitHubAccountUrl\":\"https://github.com/jpiyali/testrepository\",\"ResourceId\":\"b370eeb2-5dff-4838-8253-ec3014b2c706\"}"),
		// 				ConnectorID: to.Ptr("4C89E219-D47E-4C14-866A-018D6D634CF3"),
		// 				ID: to.Ptr("GitHub"),
		// 				IsDefault: to.Ptr(false),
		// 			},
		// 		},
		// 	},
		// }
	}
}
