package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b2965096067d6f8374b5485b0568fd36e7c9d099/specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/SessionPools_LifecycleTimed_CreateOrUpdate.json
func ExampleContainerAppsSessionPoolsClient_BeginCreateOrUpdate_createOrUpdateSessionPoolWithLifecycleTypeTimed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsSessionPoolsClient().BeginCreateOrUpdate(ctx, "rg", "testsessionpool", armappcontainers.SessionPool{
		Location: to.Ptr("East US"),
		Identity: &armappcontainers.ManagedServiceIdentity{
			Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armappcontainers.SessionPoolProperties{
			ContainerType: to.Ptr(armappcontainers.ContainerTypeCustomContainer),
			CustomContainerTemplate: &armappcontainers.CustomContainerTemplate{
				Containers: []*armappcontainers.SessionContainer{
					{
						Name: to.Ptr("testinitcontainer"),
						Args: []*string{
							to.Ptr("-c"),
							to.Ptr("while true; do echo hello; sleep 10;done")},
						Command: []*string{
							to.Ptr("/bin/sh")},
						Image: to.Ptr("repo/testcontainer:v4"),
						Resources: &armappcontainers.SessionContainerResources{
							CPU:    to.Ptr[float64](0.25),
							Memory: to.Ptr("0.5Gi"),
						},
					}},
				Ingress: &armappcontainers.SessionIngress{
					TargetPort: to.Ptr[int32](80),
				},
				RegistryCredentials: &armappcontainers.SessionRegistryCredentials{
					Identity: to.Ptr("/subscriptions/7a497526-bb8d-4816-9795-db1418a1f977/resourcegroups/test/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testSP"),
					Server:   to.Ptr("test.azurecr.io"),
				},
			},
			DynamicPoolConfiguration: &armappcontainers.DynamicPoolConfiguration{
				LifecycleConfiguration: &armappcontainers.LifecycleConfiguration{
					LifecycleType:           to.Ptr(armappcontainers.LifecycleTypeOnContainerExit),
					MaxAlivePeriodInSeconds: to.Ptr[int32](86400),
				},
			},
			EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			ManagedIdentitySettings: []*armappcontainers.ManagedIdentitySetting{
				{
					Identity:  to.Ptr("system"),
					Lifecycle: to.Ptr(armappcontainers.IdentitySettingsLifeCycleMain),
				}},
			PoolManagementType: to.Ptr(armappcontainers.PoolManagementTypeDynamic),
			ScaleConfiguration: &armappcontainers.ScaleConfiguration{
				MaxConcurrentSessions: to.Ptr[int32](500),
				ReadySessionInstances: to.Ptr[int32](100),
			},
			SessionNetworkConfiguration: &armappcontainers.SessionNetworkConfiguration{
				Status: to.Ptr(armappcontainers.SessionNetworkStatusEgressEnabled),
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
	// res.SessionPool = armappcontainers.SessionPool{
	// 	Name: to.Ptr("testsessionpool"),
	// 	Type: to.Ptr("Microsoft.App/sessionPools"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/sessionPools/testsessionpool"),
	// 	Location: to.Ptr("East US"),
	// 	Identity: &armappcontainers.ManagedServiceIdentity{
	// 		Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("bce8c037-3d10-44a4-a970-25f799611fc6"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Properties: &armappcontainers.SessionPoolProperties{
	// 		ContainerType: to.Ptr(armappcontainers.ContainerTypeCustomContainer),
	// 		CustomContainerTemplate: &armappcontainers.CustomContainerTemplate{
	// 			Containers: []*armappcontainers.SessionContainer{
	// 				{
	// 					Name: to.Ptr("testinitcontainer"),
	// 					Args: []*string{
	// 						to.Ptr("-c"),
	// 						to.Ptr("while true; do echo hello; sleep 10;done")},
	// 						Command: []*string{
	// 							to.Ptr("/bin/sh")},
	// 							Image: to.Ptr("repo/testcontainer:v4"),
	// 							Resources: &armappcontainers.SessionContainerResources{
	// 								CPU: to.Ptr[float64](0.25),
	// 								Memory: to.Ptr("0.5Gi"),
	// 							},
	// 					}},
	// 					Ingress: &armappcontainers.SessionIngress{
	// 						TargetPort: to.Ptr[int32](80),
	// 					},
	// 					RegistryCredentials: &armappcontainers.SessionRegistryCredentials{
	// 						Identity: to.Ptr("/subscriptions/7a497526-bb8d-4816-9795-db1418a1f977/resourcegroups/test/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testSP"),
	// 						Server: to.Ptr("test.azurecr.io"),
	// 					},
	// 				},
	// 				DynamicPoolConfiguration: &armappcontainers.DynamicPoolConfiguration{
	// 					LifecycleConfiguration: &armappcontainers.LifecycleConfiguration{
	// 						LifecycleType: to.Ptr(armappcontainers.LifecycleTypeTimed),
	// 						MaxAlivePeriodInSeconds: to.Ptr[int32](600),
	// 					},
	// 				},
	// 				EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
	// 				NodeCount: to.Ptr[int32](1),
	// 				PoolManagementEndpoint: to.Ptr("https://testsessionpool.agreeableriver-3d30edf1.eastus.azurecontainerapps.io"),
	// 				PoolManagementType: to.Ptr(armappcontainers.PoolManagementTypeDynamic),
	// 				ProvisioningState: to.Ptr(armappcontainers.SessionPoolProvisioningStateSucceeded),
	// 				ScaleConfiguration: &armappcontainers.ScaleConfiguration{
	// 					MaxConcurrentSessions: to.Ptr[int32](500),
	// 					ReadySessionInstances: to.Ptr[int32](100),
	// 				},
	// 				SessionNetworkConfiguration: &armappcontainers.SessionNetworkConfiguration{
	// 					Status: to.Ptr(armappcontainers.SessionNetworkStatusEgressEnabled),
	// 				},
	// 			},
	// 		}
}
