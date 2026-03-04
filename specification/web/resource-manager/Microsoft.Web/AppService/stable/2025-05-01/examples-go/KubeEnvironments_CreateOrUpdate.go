package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/KubeEnvironments_CreateOrUpdate.json
func ExampleKubeEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKubeEnvironmentsClient().BeginCreateOrUpdate(ctx, "examplerg", "testkubeenv", armappservice.KubeEnvironment{
		Location: to.Ptr("East US"),
		Properties: &armappservice.KubeEnvironmentProperties{
			StaticIP: to.Ptr("1.2.3.4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappservice.KubeEnvironmentsClientCreateOrUpdateResponse{
	// 	KubeEnvironment: &armappservice.KubeEnvironment{
	// 		Name: to.Ptr("testkubeenv"),
	// 		Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
	// 		ExtendedLocation: &armappservice.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 			Type: to.Ptr("customLocation"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/testkubeenv"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappservice.KubeEnvironmentProperties{
	// 			AksResourceID: to.Ptr("test"),
	// 			DefaultDomain: to.Ptr("testkubeenv.k4apps.io"),
	// 			InternalLoadBalancerEnabled: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
	// 			StaticIP: to.Ptr("1.2.3.4"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
