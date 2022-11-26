package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2022-09-01/examples/ContainerGroupExtensions.json
func ExampleContainerGroupsClient_BeginCreateOrUpdate_containerGroupCreateWithExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerinstance.NewContainerGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "demo", "demo1", armcontainerinstance.ContainerGroup{
		Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
			Containers: []*armcontainerinstance.Container{
				{
					Name: to.Ptr("demo1"),
					Properties: &armcontainerinstance.ContainerProperties{
						Command:              []*string{},
						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{},
						Image:                to.Ptr("nginx"),
						Ports: []*armcontainerinstance.ContainerPort{
							{
								Port: to.Ptr[int32](80),
							}},
						Resources: &armcontainerinstance.ResourceRequirements{
							Requests: &armcontainerinstance.ResourceRequests{
								CPU:        to.Ptr[float64](1),
								MemoryInGB: to.Ptr[float64](1.5),
							},
						},
					},
				}},
			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{},
			IPAddress: &armcontainerinstance.IPAddress{
				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePrivate),
				Ports: []*armcontainerinstance.Port{
					{
						Port:     to.Ptr[int32](80),
						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
					}},
			},
			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
			SubnetIDs: []*armcontainerinstance.ContainerGroupSubnetID{
				{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-rg-vnet/subnets/test-subnet"),
				}},
			Extensions: []*armcontainerinstance.DeploymentExtensionSpec{
				{
					Name: to.Ptr("kube-proxy"),
					Properties: &armcontainerinstance.DeploymentExtensionSpecProperties{
						ExtensionType: to.Ptr("kube-proxy"),
						ProtectedSettings: map[string]interface{}{
							"kubeConfig": "<kubeconfig encoded string>",
						},
						Settings: map[string]interface{}{
							"clusterCidr": "10.240.0.0/16",
							"kubeVersion": "v1.9.10",
						},
						Version: to.Ptr("1.0"),
					},
				},
				{
					Name: to.Ptr("vk-realtime-metrics"),
					Properties: &armcontainerinstance.DeploymentExtensionSpecProperties{
						ExtensionType: to.Ptr("realtime-metrics"),
						Version:       to.Ptr("1.0"),
					},
				}},
		},
		Location: to.Ptr("eastus2"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
