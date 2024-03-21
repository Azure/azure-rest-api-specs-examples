package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/JavaComponents_List_ServiceBind.json
func ExampleJavaComponentsClient_NewListPager_listJavaComponentsWithServiceBinds() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJavaComponentsClient().NewListPager("examplerg", "myenvironment", nil)
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
		// page.JavaComponentsCollection = armappcontainers.JavaComponentsCollection{
		// 	Value: []*armappcontainers.JavaComponent{
		// 		{
		// 			Name: to.Ptr("blueshark"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/javaComponents"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/blueshark"),
		// 			Properties: &armappcontainers.JavaComponentProperties{
		// 				ComponentType: to.Ptr(armappcontainers.JavaComponentTypeSpringBootAdmin),
		// 				Configurations: []*armappcontainers.JavaComponentConfigurationProperty{
		// 					{
		// 						PropertyName: to.Ptr("spring.boot.admin.ui.enable-toasts"),
		// 						Value: to.Ptr("true"),
		// 					},
		// 					{
		// 						PropertyName: to.Ptr("spring.boot.admin.monitor.status-interval"),
		// 						Value: to.Ptr("10000ms"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armappcontainers.JavaComponentProvisioningStateSucceeded),
		// 				ServiceBinds: []*armappcontainers.JavaComponentServiceBind{
		// 					{
		// 						Name: to.Ptr("yellowcat"),
		// 						ServiceID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/yellowcat"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("yellowcat"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/javaComponents"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/yellowcat"),
		// 			Properties: &armappcontainers.JavaComponentProperties{
		// 				ComponentType: to.Ptr(armappcontainers.JavaComponentTypeSpringCloudEureka),
		// 				Configurations: []*armappcontainers.JavaComponentConfigurationProperty{
		// 					{
		// 						PropertyName: to.Ptr("spring.cloud.config.server.git.uri"),
		// 						Value: to.Ptr("<GIT-URL>"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armappcontainers.JavaComponentProvisioningStateSucceeded),
		// 				ServiceBinds: []*armappcontainers.JavaComponentServiceBind{
		// 					{
		// 						Name: to.Ptr("blueshark"),
		// 						ServiceID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/blueshark"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
