package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Compute/get/ComputeInstance.json
func ExampleComputeClient_Get_getAnComputeInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputeClient().Get(ctx, "testrg123", "workspaces123", "compute123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ComputeClientGetResponse{
	// 	ComputeResource: armmachinelearning.ComputeResource{
	// 		Name: to.Ptr("compute123"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armmachinelearning.ComputeInstance{
	// 			Description: to.Ptr("some compute"),
	// 			ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
	// 			CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.0000000+00:00"); return t}()),
	// 			ModifiedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.0000000+00:00"); return t}()),
	// 			Properties: &armmachinelearning.ComputeInstanceProperties{
	// 				ApplicationSharingPolicy: to.Ptr(armmachinelearning.ApplicationSharingPolicyShared),
	// 				Applications: []*armmachinelearning.ComputeInstanceApplication{
	// 					{
	// 						DisplayName: to.Ptr("Jupyter"),
	// 						EndpointURI: to.Ptr("https://compute123.eastus2.azureml.net/jupyter"),
	// 					},
	// 				},
	// 				ComputeInstanceAuthorizationType: to.Ptr(armmachinelearning.ComputeInstanceAuthorizationTypePersonal),
	// 				ConnectivityEndpoints: &armmachinelearning.ComputeInstanceConnectivityEndpoints{
	// 					PrivateIPAddress: to.Ptr("10.0.0.1"),
	// 					PublicIPAddress: to.Ptr("10.0.0.1"),
	// 				},
	// 				CreatedBy: &armmachinelearning.ComputeInstanceCreatedBy{
	// 					UserID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					UserName: to.Ptr("foobar@microsoft.com"),
	// 					UserOrgID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 				CustomServices: []*armmachinelearning.CustomService{
	// 					{
	// 						Name: to.Ptr("rstudio-workbench"),
	// 						Docker: &armmachinelearning.Docker{
	// 							Privileged: to.Ptr(true),
	// 						},
	// 						Endpoints: []*armmachinelearning.Endpoint{
	// 							{
	// 								Name: to.Ptr("connect"),
	// 								HostIP: nil,
	// 								Published: to.Ptr[int32](4444),
	// 								Target: to.Ptr[int32](8787),
	// 								Protocol: to.Ptr(armmachinelearning.ProtocolHTTP),
	// 							},
	// 						},
	// 						EnvironmentVariables: map[string]*armmachinelearning.EnvironmentVariable{
	// 							"RSP_LICENSE": &armmachinelearning.EnvironmentVariable{
	// 								Type: to.Ptr(armmachinelearning.EnvironmentVariableTypeLocal),
	// 								Value: to.Ptr("XXXX-XXXX-XXXX-XXXX-XXXX-XXXX-XXXX"),
	// 							},
	// 						},
	// 						Image: &armmachinelearning.Image{
	// 							Type: to.Ptr(armmachinelearning.ImageTypeDocker),
	// 							Reference: to.Ptr("ghcr.io/azure/rstudio-workbench:latest"),
	// 						},
	// 						Volumes: []*armmachinelearning.VolumeDefinition{
	// 							{
	// 								Type: to.Ptr(armmachinelearning.VolumeDefinitionTypeBind),
	// 								Bind: &armmachinelearning.BindOptions{
	// 									CreateHostPath: to.Ptr(true),
	// 									Propagation: to.Ptr("test"),
	// 									Selinux: to.Ptr("test"),
	// 								},
	// 								Consistency: to.Ptr("test"),
	// 								ReadOnly: to.Ptr(true),
	// 								Source: to.Ptr("/mnt/azureuser/"),
	// 								Target: to.Ptr("/home/testuser/"),
	// 								Tmpfs: &armmachinelearning.TmpfsOptions{
	// 									Size: to.Ptr[int32](10),
	// 								},
	// 								Volume: &armmachinelearning.VolumeOptions{
	// 									Nocopy: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				EnableOSPatching: to.Ptr(true),
	// 				EnableRootAccess: to.Ptr(true),
	// 				EnableSSO: to.Ptr(true),
	// 				OSImageMetadata: &armmachinelearning.ImageMetadata{
	// 					CurrentImageVersion: to.Ptr("22.06.14"),
	// 					IsLatestOsImageVersion: to.Ptr(false),
	// 					LatestImageVersion: to.Ptr("22.07.22"),
	// 				},
	// 				PersonalComputeInstanceSettings: &armmachinelearning.PersonalComputeInstanceSettings{
	// 					AssignedUser: &armmachinelearning.AssignedUser{
	// 						ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					},
	// 				},
	// 				ReleaseQuotaOnStop: to.Ptr(true),
	// 				SSHSettings: &armmachinelearning.ComputeInstanceSSHSettings{
	// 					AdminUserName: to.Ptr("azureuser"),
	// 					SSHPort: to.Ptr[int32](22),
	// 					SSHPublicAccess: to.Ptr(armmachinelearning.SSHPublicAccessEnabled),
	// 				},
	// 				State: to.Ptr(armmachinelearning.ComputeInstanceStateRunning),
	// 				Subnet: &armmachinelearning.ResourceID{
	// 					ID: to.Ptr("test-subnet-resource-id"),
	// 				},
	// 				VMSize: to.Ptr("STANDARD_NC6"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
