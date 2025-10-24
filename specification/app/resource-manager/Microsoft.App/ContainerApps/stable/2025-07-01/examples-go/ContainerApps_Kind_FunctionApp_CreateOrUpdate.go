package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b2965096067d6f8374b5485b0568fd36e7c9d099/specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ContainerApps_Kind_FunctionApp_CreateOrUpdate.json
func ExampleContainerAppsClient_BeginCreateOrUpdate_createOrUpdateFunctionAppKind() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsClient().BeginCreateOrUpdate(ctx, "rg", "testcontainerAppFunctionKind", armappcontainers.ContainerApp{
		Location:  to.Ptr("East Us"),
		Kind:      to.Ptr(armappcontainers.KindFunctionapp),
		ManagedBy: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppFunctionKind"),
		Properties: &armappcontainers.ContainerAppProperties{
			Configuration: &armappcontainers.Configuration{
				ActiveRevisionsMode: to.Ptr(armappcontainers.ActiveRevisionsModeSingle),
				Ingress: &armappcontainers.Ingress{
					AllowInsecure: to.Ptr(false),
					External:      to.Ptr(true),
					TargetPort:    to.Ptr[int32](80),
				},
			},
			ManagedEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
			Template: &armappcontainers.Template{
				Containers: []*armappcontainers.Container{
					{
						Name: to.Ptr("function-app-container"),
						Env: []*armappcontainers.EnvironmentVar{
							{
								Name:  to.Ptr("AzureWebJobsStorage"),
								Value: to.Ptr("DefaultEndpointsProtocol=https;AccountName=mystorageaccount;AccountKey=mykey;EndpointSuffix=core.windows.net"),
							},
							{
								Name:  to.Ptr("FUNCTIONS_WORKER_RUNTIME"),
								Value: to.Ptr("dotnet"),
							},
							{
								Name:  to.Ptr("WEBSITES_ENABLE_APP_SERVICE_STORAGE"),
								Value: to.Ptr("false"),
							}},
						Image: to.Ptr("mcr.microsoft.com/azure-functions/dotnet:4"),
						Resources: &armappcontainers.ContainerResources{
							CPU:    to.Ptr[float64](0.5),
							Memory: to.Ptr("1.0Gi"),
						},
					}},
				Scale: &armappcontainers.Scale{
					CooldownPeriod:  to.Ptr[int32](300),
					MaxReplicas:     to.Ptr[int32](10),
					MinReplicas:     to.Ptr[int32](0),
					PollingInterval: to.Ptr[int32](30),
				},
			},
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
	// res.ContainerApp = armappcontainers.ContainerApp{
	// 	Name: to.Ptr("testcontainerAppFunctionKind"),
	// 	Type: to.Ptr("Microsoft.App/containerApps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerapps/testcontainerAppFunctionKind"),
	// 	Location: to.Ptr("East US"),
	// 	Identity: &armappcontainers.ManagedServiceIdentity{
	// 		Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Kind: to.Ptr(armappcontainers.KindFunctionapp),
	// 	ManagedBy: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppFunctionKind"),
	// 	Properties: &armappcontainers.ContainerAppProperties{
	// 		Configuration: &armappcontainers.Configuration{
	// 			ActiveRevisionsMode: to.Ptr(armappcontainers.ActiveRevisionsModeSingle),
	// 			Ingress: &armappcontainers.Ingress{
	// 				ExposedPort: to.Ptr[int32](0),
	// 				External: to.Ptr(true),
	// 				Fqdn: to.Ptr("testcontainerAppFunctionKind.nicefield-53acf186.eastus.azurecontainerapps-test.io"),
	// 				TargetPort: to.Ptr[int32](80),
	// 				Traffic: []*armappcontainers.TrafficWeight{
	// 					{
	// 						LatestRevision: to.Ptr(true),
	// 						Weight: to.Ptr[int32](100),
	// 				}},
	// 				Transport: to.Ptr(armappcontainers.IngressTransportMethodAuto),
	// 			},
	// 		},
	// 		EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
	// 		EventStreamEndpoint: to.Ptr("https://azurecontainerapps-test.dev/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/containerApps/testcontainerAppFunctionKind/eventstream"),
	// 		LatestReadyRevisionName: to.Ptr("testcontainerAppFunctionKind--abc123"),
	// 		LatestRevisionFqdn: to.Ptr("testcontainerAppFunctionKind--abc123.nicefield-53acf186.eastus.azurecontainerapps-test.io"),
	// 		LatestRevisionName: to.Ptr("testcontainerAppFunctionKind--abc123"),
	// 		ManagedEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ContainerAppProvisioningStateSucceeded),
	// 		RunningStatus: to.Ptr(armappcontainers.ContainerAppRunningStatusRunning),
	// 		Template: &armappcontainers.Template{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("function-app-container"),
	// 					Env: []*armappcontainers.EnvironmentVar{
	// 						{
	// 							Name: to.Ptr("AzureWebJobsStorage"),
	// 							Value: to.Ptr("DefaultEndpointsProtocol=https;AccountName=mystorageaccount;AccountKey=mykey;EndpointSuffix=core.windows.net"),
	// 						},
	// 						{
	// 							Name: to.Ptr("FUNCTIONS_WORKER_RUNTIME"),
	// 							Value: to.Ptr("dotnet"),
	// 						},
	// 						{
	// 							Name: to.Ptr("WEBSITES_ENABLE_APP_SERVICE_STORAGE"),
	// 							Value: to.Ptr("false"),
	// 					}},
	// 					Image: to.Ptr("mcr.microsoft.com/azure-functions/dotnet:4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.5),
	// 						EphemeralStorage: to.Ptr("2Gi"),
	// 						Memory: to.Ptr("1Gi"),
	// 					},
	// 			}},
	// 			RevisionSuffix: to.Ptr(""),
	// 			Scale: &armappcontainers.Scale{
	// 				CooldownPeriod: to.Ptr[int32](300),
	// 				MaxReplicas: to.Ptr[int32](10),
	// 				MinReplicas: to.Ptr[int32](0),
	// 				PollingInterval: to.Ptr[int32](30),
	// 			},
	// 		},
	// 	},
	// }
}
