package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/JavaComponents_Get.json
func ExampleJavaComponentsClient_Get_getJavaComponent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJavaComponentsClient().Get(ctx, "examplerg", "myenvironment", "myjavacomponent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JavaComponent = armappcontainers.JavaComponent{
	// 	Name: to.Ptr("myjavacomponent"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments/javaComponents"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/myjavacomponent"),
	// 	Properties: &armappcontainers.SpringBootAdminComponent{
	// 		ComponentType: to.Ptr(armappcontainers.JavaComponentTypeSpringBootAdmin),
	// 		Configurations: []*armappcontainers.JavaComponentConfigurationProperty{
	// 			{
	// 				PropertyName: to.Ptr("spring.boot.admin.ui.enable-toasts"),
	// 				Value: to.Ptr("true"),
	// 			},
	// 			{
	// 				PropertyName: to.Ptr("spring.boot.admin.monitor.status-interval"),
	// 				Value: to.Ptr("10000ms"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armappcontainers.JavaComponentProvisioningStateSucceeded),
	// 		Scale: &armappcontainers.JavaComponentPropertiesScale{
	// 			MaxReplicas: to.Ptr[int32](1),
	// 			MinReplicas: to.Ptr[int32](1),
	// 		},
	// 		Ingress: &armappcontainers.JavaComponentIngress{
	// 			Fqdn: to.Ptr("myjavacomponent.myenvironment.test.net"),
	// 		},
	// 	},
	// }
}
