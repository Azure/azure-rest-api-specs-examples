package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/createOrUpdate/ComputeInstance.json
func ExampleComputeClient_BeginCreateOrUpdate_createAnComputeInstanceCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputeClient().BeginCreateOrUpdate(ctx, "testrg123", "workspaces123", "compute123", armmachinelearning.ComputeResource{
		Properties: &armmachinelearning.ComputeInstance{
			ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
			Properties: &armmachinelearning.ComputeInstanceProperties{
				ApplicationSharingPolicy:         to.Ptr(armmachinelearning.ApplicationSharingPolicyPersonal),
				ComputeInstanceAuthorizationType: to.Ptr(armmachinelearning.ComputeInstanceAuthorizationTypePersonal),
				CustomServices: []*armmachinelearning.CustomService{
					{
						Name: to.Ptr("rstudio"),
						Docker: &armmachinelearning.Docker{
							Privileged: to.Ptr(true),
						},
						Endpoints: []*armmachinelearning.Endpoint{
							{
								Name:      to.Ptr("connect"),
								Published: to.Ptr[int32](8787),
								Target:    to.Ptr[int32](8787),
								Protocol:  to.Ptr(armmachinelearning.ProtocolHTTP),
							}},
						EnvironmentVariables: map[string]*armmachinelearning.EnvironmentVariable{
							"test_variable": {
								Type:  to.Ptr(armmachinelearning.EnvironmentVariableTypeLocal),
								Value: to.Ptr("test_value"),
							},
						},
						Image: &armmachinelearning.Image{
							Type:      to.Ptr(armmachinelearning.ImageTypeDocker),
							Reference: to.Ptr("ghcr.io/azure/rocker-rstudio-ml-verse:latest"),
						},
						Volumes: []*armmachinelearning.VolumeDefinition{
							{
								Type:     to.Ptr(armmachinelearning.VolumeDefinitionTypeBind),
								ReadOnly: to.Ptr(false),
								Source:   to.Ptr("/home/azureuser/cloudfiles"),
								Target:   to.Ptr("/home/azureuser/cloudfiles"),
							}},
					}},
				PersonalComputeInstanceSettings: &armmachinelearning.PersonalComputeInstanceSettings{
					AssignedUser: &armmachinelearning.AssignedUser{
						ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					},
				},
				SSHSettings: &armmachinelearning.ComputeInstanceSSHSettings{
					SSHPublicAccess: to.Ptr(armmachinelearning.SSHPublicAccessDisabled),
				},
				Subnet: &armmachinelearning.ResourceID{
					ID: to.Ptr("test-subnet-resource-id"),
				},
				VMSize: to.Ptr("STANDARD_NC6"),
			},
		},
		Location: to.Ptr("eastus"),
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
	// res.ComputeResource = armmachinelearning.ComputeResource{
	// 	Properties: &armmachinelearning.ComputeInstance{
	// 		ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
	// 		ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
	// 	},
	// 	Name: to.Ptr("compute123"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
	// 	Location: to.Ptr("eastus"),
	// }
}
