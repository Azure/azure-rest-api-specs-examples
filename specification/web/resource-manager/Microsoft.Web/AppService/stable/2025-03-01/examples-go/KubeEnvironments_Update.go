package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/KubeEnvironments_Update.json
func ExampleKubeEnvironmentsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKubeEnvironmentsClient().Update(ctx, "examplerg", "testkubeenv", armappservice.KubeEnvironmentPatchResource{
		Properties: &armappservice.KubeEnvironmentPatchResourceProperties{
			StaticIP: to.Ptr("1.2.3.4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubeEnvironment = armappservice.KubeEnvironment{
	// 	Name: to.Ptr("testkubeenv"),
	// 	Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/testkubeenv"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armappservice.KubeEnvironmentProperties{
	// 		AksResourceID: to.Ptr("test"),
	// 		DefaultDomain: to.Ptr("testkubeenv.k4apps.io"),
	// 		InternalLoadBalancerEnabled: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 	},
	// }
}
