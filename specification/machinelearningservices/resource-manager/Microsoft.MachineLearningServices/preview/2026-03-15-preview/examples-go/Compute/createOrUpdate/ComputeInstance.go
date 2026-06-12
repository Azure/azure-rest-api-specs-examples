package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Compute/createOrUpdate/ComputeInstance.json
func ExampleComputeClient_BeginCreateOrUpdate_createAnComputeInstanceCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputeClient().BeginCreateOrUpdate(ctx, "testrg123", "workspaces123", "compute123", armmachinelearning.ComputeResource{
		Location: to.Ptr("eastus"),
		Properties: &armmachinelearning.ComputeInstance{
			ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
			Properties: &armmachinelearning.ComputeInstanceProperties{
				ApplicationSharingPolicy: to.Ptr(armmachinelearning.ApplicationSharingPolicyPersonal),
				AutologgerSettings: &armmachinelearning.ComputeInstanceAutologgerSettings{
					MlflowAutologger: to.Ptr(armmachinelearning.MlflowAutologgerEnabled),
				},
				ComputeInstanceAuthorizationType: to.Ptr(armmachinelearning.ComputeInstanceAuthorizationTypePersonal),
				CustomServices: []*armmachinelearning.CustomService{
					{
						Name: to.Ptr("rstudio-workbench"),
						Docker: &armmachinelearning.Docker{
							Privileged: to.Ptr(true),
						},
						Endpoints: []*armmachinelearning.Endpoint{
							{
								Name:      to.Ptr("connect"),
								HostIP:    nil,
								Published: to.Ptr[int32](4444),
								Target:    to.Ptr[int32](8787),
								Protocol:  to.Ptr(armmachinelearning.ProtocolHTTP),
							},
						},
						EnvironmentVariables: map[string]*armmachinelearning.EnvironmentVariable{
							"RSP_LICENSE": {
								Type:  to.Ptr(armmachinelearning.EnvironmentVariableTypeLocal),
								Value: to.Ptr("XXXX-XXXX-XXXX-XXXX-XXXX-XXXX-XXXX"),
							},
						},
						Image: &armmachinelearning.Image{
							Type:      to.Ptr(armmachinelearning.ImageTypeDocker),
							Reference: to.Ptr("ghcr.io/azure/rstudio-workbench:latest"),
						},
						Kernel: &armmachinelearning.JupyterKernelConfig{
							Argv: []*string{
								to.Ptr("option1"),
								to.Ptr("option2"),
								to.Ptr("option3"),
							},
							DisplayName: to.Ptr("TestKernel"),
							Language:    to.Ptr("python"),
						},
						Volumes: []*armmachinelearning.VolumeDefinition{
							{
								Type:     to.Ptr(armmachinelearning.VolumeDefinitionTypeBind),
								ReadOnly: to.Ptr(true),
								Source:   to.Ptr("/mnt/azureuser/"),
								Target:   to.Ptr("/home/testuser/"),
							},
						},
					},
				},
				EnableOSPatching: to.Ptr(true),
				EnableRootAccess: to.Ptr(true),
				EnableSSO:        to.Ptr(true),
				PersonalComputeInstanceSettings: &armmachinelearning.PersonalComputeInstanceSettings{
					AssignedUser: &armmachinelearning.AssignedUser{
						ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					},
				},
				ReleaseQuotaOnStop: to.Ptr(true),
				SSHSettings: &armmachinelearning.ComputeInstanceSSHSettings{
					SSHPublicAccess: to.Ptr(armmachinelearning.SSHPublicAccessDisabled),
				},
				Subnet: &armmachinelearning.ResourceID{
					ID: to.Ptr("test-subnet-resource-id"),
				},
				VMSize: to.Ptr("STANDARD_NC6"),
			},
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
	// res = armmachinelearning.ComputeClientCreateOrUpdateResponse{
	// 	ComputeResource: armmachinelearning.ComputeResource{
	// 		Name: to.Ptr("compute123"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmachinelearning.ComputeInstance{
	// 			ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
	// 			ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateCreating),
	// 		},
	// 	},
	// }
}
