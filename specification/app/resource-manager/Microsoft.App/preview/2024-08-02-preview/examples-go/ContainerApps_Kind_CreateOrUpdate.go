package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerApps_Kind_CreateOrUpdate.json
func ExampleContainerAppsClient_BeginCreateOrUpdate_createOrUpdateAppKind() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsClient().BeginCreateOrUpdate(ctx, "rg", "testcontainerAppKind", armappcontainers.ContainerApp{
		Location:  to.Ptr("East Us"),
		Kind:      to.Ptr(armappcontainers.KindWorkflowapp),
		ManagedBy: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppKind"),
		Properties: &armappcontainers.ContainerAppProperties{
			Configuration: &armappcontainers.Configuration{
				ActiveRevisionsMode: to.Ptr(armappcontainers.ActiveRevisionsModeSingle),
				Ingress: &armappcontainers.Ingress{
					AllowInsecure: to.Ptr(true),
					External:      to.Ptr(true),
					TargetPort:    to.Ptr[int32](80),
				},
			},
			ManagedEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
			Template: &armappcontainers.Template{
				Containers: []*armappcontainers.Container{
					{
						Name:  to.Ptr("logicapps-container"),
						Image: to.Ptr("default/logicapps-base:latest"),
						Resources: &armappcontainers.ContainerResources{
							CPU:    to.Ptr[float64](1),
							Memory: to.Ptr("2.0Gi"),
						},
					}},
				Scale: &armappcontainers.Scale{
					CooldownPeriod:  to.Ptr[int32](350),
					MaxReplicas:     to.Ptr[int32](30),
					MinReplicas:     to.Ptr[int32](1),
					PollingInterval: to.Ptr[int32](35),
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
	// 	Name: to.Ptr("testcontainerAppKind"),
	// 	Type: to.Ptr("Microsoft.App/containerApps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/containerapps/testcontainerAppKind"),
	// 	Location: to.Ptr("East US"),
	// 	Identity: &armappcontainers.ManagedServiceIdentity{
	// 		Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Kind: to.Ptr(armappcontainers.KindWorkflowapp),
	// 	ManagedBy: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppKind"),
	// 	Properties: &armappcontainers.ContainerAppProperties{
	// 		Configuration: &armappcontainers.Configuration{
	// 			ActiveRevisionsMode: to.Ptr(armappcontainers.ActiveRevisionsModeSingle),
	// 			Ingress: &armappcontainers.Ingress{
	// 				AllowInsecure: to.Ptr(true),
	// 				ExposedPort: to.Ptr[int32](0),
	// 				External: to.Ptr(true),
	// 				Fqdn: to.Ptr("testcontainerAppKind.nicefield-53acf186.eastus.azurecontainerapps-test.io"),
	// 				TargetPort: to.Ptr[int32](80),
	// 				Traffic: []*armappcontainers.TrafficWeight{
	// 					{
	// 						LatestRevision: to.Ptr(true),
	// 						Weight: to.Ptr[int32](100),
	// 				}},
	// 				Transport: to.Ptr(armappcontainers.IngressTransportMethodAuto),
	// 			},
	// 			MaxInactiveRevisions: to.Ptr[int32](100),
	// 		},
	// 		EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
	// 		EventStreamEndpoint: to.Ptr("https://azurecontainerapps-test.dev/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/containerApps/testcontainerAppKind/eventstream"),
	// 		LatestReadyRevisionName: to.Ptr("testcontainerAppKind--2rltv14"),
	// 		LatestRevisionFqdn: to.Ptr("testcontainerAppKind--2rltv14.nicefield-53acf186.eastus.azurecontainerapps-test.io"),
	// 		LatestRevisionName: to.Ptr("testcontainerAppKind--2rltv14"),
	// 		ManagedEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ContainerAppProvisioningStateSucceeded),
	// 		Template: &armappcontainers.Template{
	// 			Containers: []*armappcontainers.Container{
	// 				{
	// 					Name: to.Ptr("logicapps-container"),
	// 					Image: to.Ptr("default/logicapps-base:latest"),
	// 					ImageType: to.Ptr(armappcontainers.ImageTypeContainerImage),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](1),
	// 						EphemeralStorage: to.Ptr("4Gi"),
	// 						Memory: to.Ptr("2Gi"),
	// 					},
	// 			}},
	// 			RevisionSuffix: to.Ptr(""),
	// 			Scale: &armappcontainers.Scale{
	// 				CooldownPeriod: to.Ptr[int32](350),
	// 				MaxReplicas: to.Ptr[int32](30),
	// 				MinReplicas: to.Ptr[int32](1),
	// 				PollingInterval: to.Ptr[int32](35),
	// 			},
	// 		},
	// 	},
	// }
}
