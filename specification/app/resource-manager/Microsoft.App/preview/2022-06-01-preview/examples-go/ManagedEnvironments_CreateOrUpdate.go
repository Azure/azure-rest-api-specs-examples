package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/212686c8383679e034b19143e13cbeb5a40ab454/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ManagedEnvironments_CreateOrUpdate.json
func ExampleManagedEnvironmentsClient_BeginCreateOrUpdate() {
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
				CertificatePassword: []byte("private key password"),
				CertificateValue:    []byte("PFX-or-PEM-blob"),
				DNSSuffix:           to.Ptr("www.my-name.com"),
			},
			DaprAIConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			VnetConfiguration: &armappcontainers.VnetConfiguration{
				OutboundSettings: &armappcontainers.ManagedEnvironmentOutboundSettings{
					OutBoundType:              to.Ptr(armappcontainers.ManagedEnvironmentOutBoundTypeUserDefinedRouting),
					VirtualNetworkApplianceIP: to.Ptr("192.168.1.20"),
				},
			},
			WorkloadProfiles: []*armappcontainers.WorkloadProfile{
				{
					MaximumCount:        to.Ptr[int32](12),
					MinimumCount:        to.Ptr[int32](3),
					WorkloadProfileType: to.Ptr("GeneralPurpose"),
				},
				{
					MaximumCount:        to.Ptr[int32](6),
					MinimumCount:        to.Ptr[int32](3),
					WorkloadProfileType: to.Ptr("MemoryOptimized"),
				},
				{
					MaximumCount:        to.Ptr[int32](6),
					MinimumCount:        to.Ptr[int32](3),
					WorkloadProfileType: to.Ptr("ComputeOptimized"),
				}},
			ZoneRedundant: to.Ptr(true),
		},
		SKU: &armappcontainers.EnvironmentSKUProperties{
			Name: to.Ptr(armappcontainers.SKUNamePremium),
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
	// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
	// 			SubjectName: to.Ptr("CN=www.my-name.com"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 		},
	// 		DefaultDomain: to.Ptr("testcontainerenv.k4apps.io"),
	// 		EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 		ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 		VnetConfiguration: &armappcontainers.VnetConfiguration{
	// 			OutboundSettings: &armappcontainers.ManagedEnvironmentOutboundSettings{
	// 				OutBoundType: to.Ptr(armappcontainers.ManagedEnvironmentOutBoundTypeUserDefinedRouting),
	// 				VirtualNetworkApplianceIP: to.Ptr("192.168.1.20"),
	// 			},
	// 		},
	// 		WorkloadProfiles: []*armappcontainers.WorkloadProfile{
	// 			{
	// 				MaximumCount: to.Ptr[int32](12),
	// 				MinimumCount: to.Ptr[int32](3),
	// 				WorkloadProfileType: to.Ptr("GeneralPurpose"),
	// 			},
	// 			{
	// 				MaximumCount: to.Ptr[int32](6),
	// 				MinimumCount: to.Ptr[int32](3),
	// 				WorkloadProfileType: to.Ptr("MemoryOptimized"),
	// 			},
	// 			{
	// 				MaximumCount: to.Ptr[int32](6),
	// 				MinimumCount: to.Ptr[int32](3),
	// 				WorkloadProfileType: to.Ptr("ComputeOptimized"),
	// 		}},
	// 		ZoneRedundant: to.Ptr(true),
	// 	},
	// 	SKU: &armappcontainers.EnvironmentSKUProperties{
	// 		Name: to.Ptr(armappcontainers.SKUNamePremium),
	// 	},
	// }
}
