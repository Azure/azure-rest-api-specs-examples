package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9a65970ef1837c0af1800c906aa365ba95871b26/specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_Get.json
func ExampleManagedEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentsClient().Get(ctx, "examplerg", "jlaw-demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedEnvironment = armappcontainers.ManagedEnvironment{
	// 	Name: to.Ptr("jlaw-demo1"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/jlaw-demo1"),
	// 	Location: to.Ptr("North Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armappcontainers.ManagedEnvironmentProperties{
	// 		DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
	// 		ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("20.42.33.145"),
	// 		ZoneRedundant: to.Ptr(true),
	// 	},
	// }
}
