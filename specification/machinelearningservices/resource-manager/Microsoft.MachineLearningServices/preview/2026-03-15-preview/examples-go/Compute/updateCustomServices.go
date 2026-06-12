package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Compute/updateCustomServices.json
func ExampleComputeClient_UpdateCustomServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputeClient().UpdateCustomServices(ctx, "testrg123", "workspaces123", "compute123", []*armmachinelearning.CustomService{
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
			Volumes: []*armmachinelearning.VolumeDefinition{
				{
					Type:     to.Ptr(armmachinelearning.VolumeDefinitionTypeBind),
					ReadOnly: to.Ptr(true),
					Source:   to.Ptr("/mnt/azureuser/"),
					Target:   to.Ptr("/home/testuser/"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ComputeClientUpdateCustomServicesResponse{
	// }
}
