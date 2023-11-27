package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedEnvironments_CreateOrUpdate.json
func ExampleManagedEnvironmentsClient_BeginCreateOrUpdate_createEnvironments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedEnvironmentsClient().BeginCreateOrUpdate(ctx, "examplerg", "testcontainerenv", armappcontainers.ManagedEnvironment{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ManagedEnvironmentProperties{
			AppLogsConfiguration: &armappcontainers.AppLogsConfiguration{
				LogAnalyticsConfiguration: &armappcontainers.LogAnalyticsConfiguration{
					CustomerID: to.Ptr("string"),
					SharedKey:  to.Ptr("string"),
				},
			},
			CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
				CertificatePassword: to.Ptr("1234"),
				CertificateValue:    []byte("Y2VydA=="),
				DNSSuffix:           to.Ptr("www.my-name.com"),
			},
			DaprAIConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			PeerAuthentication: &armappcontainers.ManagedEnvironmentPropertiesPeerAuthentication{
				Mtls: &armappcontainers.Mtls{
					Enabled: to.Ptr(true),
				},
			},
			VnetConfiguration: &armappcontainers.VnetConfiguration{
				InfrastructureSubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"),
			},
			WorkloadProfiles: []*armappcontainers.WorkloadProfile{
				{
					Name:                to.Ptr("My-GP-01"),
					MaximumCount:        to.Ptr[int32](12),
					MinimumCount:        to.Ptr[int32](3),
					WorkloadProfileType: to.Ptr("GeneralPurpose"),
				},
				{
					Name:                to.Ptr("My-MO-01"),
					MaximumCount:        to.Ptr[int32](6),
					MinimumCount:        to.Ptr[int32](3),
					WorkloadProfileType: to.Ptr("MemoryOptimized"),
				},
				{
					Name:                to.Ptr("My-CO-01"),
					MaximumCount:        to.Ptr[int32](6),
					MinimumCount:        to.Ptr[int32](3),
					WorkloadProfileType: to.Ptr("ComputeOptimized"),
				},
				{
					Name:                to.Ptr("My-consumption-01"),
					WorkloadProfileType: to.Ptr("Consumption"),
				}},
			ZoneRedundant: to.Ptr(true),
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
	// res.ManagedEnvironment = armappcontainers.ManagedEnvironment{
	// 	Name: to.Ptr("testcontainerenv"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ManagedEnvironmentProperties{
	// 		AppLogsConfiguration: &armappcontainers.AppLogsConfiguration{
	// 			LogAnalyticsConfiguration: &armappcontainers.LogAnalyticsConfiguration{
	// 				CustomerID: to.Ptr("string"),
	// 			},
	// 		},
	// 		CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 			CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 			DNSSuffix: to.Ptr("www.my-name.com"),
	// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00.000Z"); return t}()),
	// 			SubjectName: to.Ptr("CN=www.my-name.com"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 		},
	// 		DefaultDomain: to.Ptr("testcontainerenv.k4apps.io"),
	// 		EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 		InfrastructureResourceGroup: to.Ptr("capp-svc-testcontainerenv-eastus"),
	// 		PeerAuthentication: &armappcontainers.ManagedEnvironmentPropertiesPeerAuthentication{
	// 			Mtls: &armappcontainers.Mtls{
	// 				Enabled: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 		VnetConfiguration: &armappcontainers.VnetConfiguration{
	// 			InfrastructureSubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"),
	// 		},
	// 		WorkloadProfiles: []*armappcontainers.WorkloadProfile{
	// 			{
	// 				Name: to.Ptr("My-GP-01"),
	// 				MaximumCount: to.Ptr[int32](12),
	// 				MinimumCount: to.Ptr[int32](3),
	// 				WorkloadProfileType: to.Ptr("GeneralPurpose"),
	// 			},
	// 			{
	// 				Name: to.Ptr("My-MO-01"),
	// 				MaximumCount: to.Ptr[int32](6),
	// 				MinimumCount: to.Ptr[int32](3),
	// 				WorkloadProfileType: to.Ptr("MemoryOptimized"),
	// 			},
	// 			{
	// 				Name: to.Ptr("My-CO-01"),
	// 				MaximumCount: to.Ptr[int32](6),
	// 				MinimumCount: to.Ptr[int32](3),
	// 				WorkloadProfileType: to.Ptr("ComputeOptimized"),
	// 			},
	// 			{
	// 				Name: to.Ptr("My-consumption-01"),
	// 				WorkloadProfileType: to.Ptr("Consumption"),
	// 		}},
	// 		ZoneRedundant: to.Ptr(true),
	// 	},
	// }
}
