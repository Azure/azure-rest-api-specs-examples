package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ManagedEnvironments_ListBySubscription.json
func ExampleManagedEnvironmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedEnvironmentsClient().NewListBySubscriptionPager(nil)
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
		// page.ManagedEnvironmentsCollection = armappcontainers.ManagedEnvironmentsCollection{
		// 	Value: []*armappcontainers.ManagedEnvironment{
		// 		{
		// 			Name: to.Ptr("jlaw-demo1"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/jlaw-demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappcontainers.ManagedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
		// 				EventStreamEndpoint: to.Ptr("testEndpoint"),
		// 				InfrastructureResourceGroup: to.Ptr("capp-svc-jlaw-demo1-northcentralus"),
		// 				ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 				VnetConfiguration: &armappcontainers.VnetConfiguration{
		// 					InfrastructureSubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"),
		// 				},
		// 				WorkloadProfiles: []*armappcontainers.WorkloadProfile{
		// 					{
		// 						Name: to.Ptr("My-GP-01"),
		// 						MaximumCount: to.Ptr[int32](12),
		// 						MinimumCount: to.Ptr[int32](3),
		// 						WorkloadProfileType: to.Ptr("GeneralPurpose"),
		// 					},
		// 					{
		// 						Name: to.Ptr("My-MO-01"),
		// 						MaximumCount: to.Ptr[int32](6),
		// 						MinimumCount: to.Ptr[int32](3),
		// 						WorkloadProfileType: to.Ptr("MemoryOptimized"),
		// 					},
		// 					{
		// 						Name: to.Ptr("My-CO-01"),
		// 						MaximumCount: to.Ptr[int32](6),
		// 						MinimumCount: to.Ptr[int32](3),
		// 						WorkloadProfileType: to.Ptr("ComputeOptimized"),
		// 					},
		// 					{
		// 						Name: to.Ptr("My-consumption-01"),
		// 						WorkloadProfileType: to.Ptr("Consumption"),
		// 				}},
		// 				ZoneRedundant: to.Ptr(true),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("demo1"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/DemoRG/providers/Microsoft.App/managedEnvironments/demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappcontainers.ManagedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name2.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name2.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("demo1.k4apps.io"),
		// 				EventStreamEndpoint: to.Ptr("testEndpoint"),
		// 				InfrastructureResourceGroup: to.Ptr("capp-svc-demo1-northcentralus"),
		// 				ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 				VnetConfiguration: &armappcontainers.VnetConfiguration{
		// 					InfrastructureSubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"),
		// 				},
		// 				ZoneRedundant: to.Ptr(true),
		// 			},
		// 	}},
		// }
	}
}
