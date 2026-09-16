package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/ManagedEnvironments_CreateOrUpdate.json
func ExampleManagedEnvironmentsClient_BeginCreateOrUpdate_createEnvironments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedEnvironmentsClient().BeginCreateOrUpdate(ctx, "examplerg", "testcontainerenv", armappcontainers.ManagedEnvironment{
		Identity: &armappcontainers.ManagedServiceIdentity{
			Type: to.Ptr(armappcontainers.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
			UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": {},
			},
		},
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ManagedEnvironmentProperties{
			AppLogsConfiguration: &armappcontainers.AppLogsConfiguration{
				LogAnalyticsConfiguration: &armappcontainers.LogAnalyticsConfiguration{
					CustomerID: to.Ptr("string"),
					SharedKey:  to.Ptr("string"),
				},
			},
			AppInsightsConfiguration: &armappcontainers.AppInsightsConfiguration{
				ConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://eastus-8.in.applicationinsights.azure.com/;LiveEndpoint=https://eastus.livediagnostics.monitor.azure.com/"),
			},
			OpenTelemetryConfiguration: &armappcontainers.OpenTelemetryConfiguration{
				DestinationsConfiguration: &armappcontainers.DestinationsConfiguration{
					DataDogConfiguration: &armappcontainers.DataDogConfiguration{
						Site: to.Ptr("string"),
						Key:  to.Ptr("000000000000000000000000"),
					},
					OtlpConfigurations: []*armappcontainers.OtlpConfiguration{
						{
							Name:     to.Ptr("dashboard"),
							Endpoint: to.Ptr("dashboard.k8s.region.azurecontainerapps.io:80"),
							Insecure: to.Ptr(true),
							Headers: []*armappcontainers.Header{
								{
									Key:   to.Ptr("api-key"),
									Value: to.Ptr("xxxxxxxxxxx"),
								},
							},
						},
					},
				},
				TracesConfiguration: &armappcontainers.TracesConfiguration{
					IncludeDapr: to.Ptr(true),
					Destinations: []*string{
						to.Ptr("appInsights"),
					},
				},
				LogsConfiguration: &armappcontainers.LogsConfiguration{
					Destinations: []*string{
						to.Ptr("appInsights"),
					},
				},
				MetricsConfiguration: &armappcontainers.MetricsConfiguration{
					IncludeKeda: to.Ptr(true),
					Destinations: []*string{
						to.Ptr("dataDog"),
					},
				},
			},
			CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
				CertificatePassword: to.Ptr("1234"),
				CertificateValue:    []byte("Y2VydA=="),
				DNSSuffix:           to.Ptr("www.my-name.com"),
			},
			DaprAIConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			IngressConfiguration: &armappcontainers.IngressConfiguration{
				HeaderCountLimit:              to.Ptr[int32](30),
				RequestIdleTimeout:            to.Ptr[int32](5),
				TerminationGracePeriodSeconds: to.Ptr[int32](3600),
				WorkloadProfileName:           to.Ptr("My-CO-01"),
			},
			PeerAuthentication: &armappcontainers.ManagedEnvironmentPropertiesPeerAuthentication{
				Mtls: &armappcontainers.Mtls{
					Enabled: to.Ptr(true),
				},
			},
			PeerTrafficConfiguration: &armappcontainers.ManagedEnvironmentPropertiesPeerTrafficConfiguration{
				Encryption: &armappcontainers.ManagedEnvironmentPropertiesPeerTrafficConfigurationEncryption{
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
				},
			},
			ZoneRedundant: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ManagedEnvironmentsClientCreateOrUpdateResponse{
	// 	ManagedEnvironment: armappcontainers.ManagedEnvironment{
	// 		Name: to.Ptr("testcontainerenv"),
	// 		Type: to.Ptr("Microsoft.App/managedEnvironments"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv"),
	// 		Identity: &armappcontainers.ManagedServiceIdentity{
	// 			Type: to.Ptr(armappcontainers.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": &armappcontainers.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappcontainers.ManagedEnvironmentProperties{
	// 			AppLogsConfiguration: &armappcontainers.AppLogsConfiguration{
	// 				LogAnalyticsConfiguration: &armappcontainers.LogAnalyticsConfiguration{
	// 					CustomerID: to.Ptr("string"),
	// 				},
	// 			},
	// 			OpenTelemetryConfiguration: &armappcontainers.OpenTelemetryConfiguration{
	// 				DestinationsConfiguration: &armappcontainers.DestinationsConfiguration{
	// 					DataDogConfiguration: &armappcontainers.DataDogConfiguration{
	// 						Site: to.Ptr("string"),
	// 					},
	// 					OtlpConfigurations: []*armappcontainers.OtlpConfiguration{
	// 						{
	// 							Name: to.Ptr("dashboard"),
	// 							Endpoint: to.Ptr("dashboard.k8s.region.azurecontainerapps.io:80"),
	// 							Insecure: to.Ptr(true),
	// 							Headers: []*armappcontainers.Header{
	// 								{
	// 									Key: to.Ptr("api-key"),
	// 									Value: to.Ptr("xxxxxxxxxxx"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				TracesConfiguration: &armappcontainers.TracesConfiguration{
	// 					IncludeDapr: to.Ptr(true),
	// 					Destinations: []*string{
	// 						to.Ptr("appInsights"),
	// 					},
	// 				},
	// 				LogsConfiguration: &armappcontainers.LogsConfiguration{
	// 					Destinations: []*string{
	// 						to.Ptr("appInsights"),
	// 					},
	// 				},
	// 				MetricsConfiguration: &armappcontainers.MetricsConfiguration{
	// 					IncludeKeda: to.Ptr(true),
	// 					Destinations: []*string{
	// 						to.Ptr("dataDog"),
	// 					},
	// 				},
	// 			},
	// 			CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 				CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 				DNSSuffix: to.Ptr("www.my-name.com"),
	// 				ExpirationDate: to.Ptr(time.Date(2022, time.November, 6, 4, 0, 0, 0, time.UTC)),
	// 				SubjectName: to.Ptr("CN=www.my-name.com"),
	// 				Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 			},
	// 			DefaultDomain: to.Ptr("testcontainerenv.k4apps.io"),
	// 			EventStreamEndpoint: to.Ptr("testEndpoint"),
	// 			InfrastructureResourceGroup: to.Ptr("capp-svc-testcontainerenv-eastus"),
	// 			IngressConfiguration: &armappcontainers.IngressConfiguration{
	// 				HeaderCountLimit: to.Ptr[int32](30),
	// 				RequestIdleTimeout: to.Ptr[int32](5),
	// 				TerminationGracePeriodSeconds: to.Ptr[int32](3600),
	// 				WorkloadProfileName: to.Ptr("My-CO-01"),
	// 			},
	// 			EnvironmentMode: to.Ptr(armappcontainers.ManagedEnvironmentModeWorkloadProfiles),
	// 			PeerAuthentication: &armappcontainers.ManagedEnvironmentPropertiesPeerAuthentication{
	// 				Mtls: &armappcontainers.Mtls{
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
	// 			StaticIP: to.Ptr("1.2.3.4"),
	// 			VnetConfiguration: &armappcontainers.VnetConfiguration{
	// 				InfrastructureSubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"),
	// 			},
	// 			WorkloadProfiles: []*armappcontainers.WorkloadProfile{
	// 				{
	// 					Name: to.Ptr("My-GP-01"),
	// 					MaximumCount: to.Ptr[int32](12),
	// 					MinimumCount: to.Ptr[int32](3),
	// 					WorkloadProfileType: to.Ptr("GeneralPurpose"),
	// 				},
	// 				{
	// 					Name: to.Ptr("My-MO-01"),
	// 					MaximumCount: to.Ptr[int32](6),
	// 					MinimumCount: to.Ptr[int32](3),
	// 					WorkloadProfileType: to.Ptr("MemoryOptimized"),
	// 				},
	// 				{
	// 					Name: to.Ptr("My-CO-01"),
	// 					MaximumCount: to.Ptr[int32](6),
	// 					MinimumCount: to.Ptr[int32](3),
	// 					WorkloadProfileType: to.Ptr("ComputeOptimized"),
	// 				},
	// 				{
	// 					Name: to.Ptr("My-consumption-01"),
	// 					WorkloadProfileType: to.Ptr("Consumption"),
	// 				},
	// 			},
	// 			ZoneRedundant: to.Ptr(true),
	// 		},
	// 	},
	// }
}
