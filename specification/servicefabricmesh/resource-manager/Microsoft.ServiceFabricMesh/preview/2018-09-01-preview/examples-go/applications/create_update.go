package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/create_update.json
func ExampleApplicationClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabricmesh.NewApplicationClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"sbz_demo",
		"sampleApplication",
		armservicefabricmesh.ApplicationResourceDescription{
			Location: to.Ptr("EastUS"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.ApplicationResourceProperties{
				Description: to.Ptr("Service Fabric Mesh sample application."),
				Services: []*armservicefabricmesh.ServiceResourceDescription{
					{
						Name: to.Ptr("helloWorldService"),
						Properties: &armservicefabricmesh.ServiceResourceProperties{
							Description:  to.Ptr("SeaBreeze Hello World Service."),
							ReplicaCount: to.Ptr[int32](1),
							CodePackages: []*armservicefabricmesh.ContainerCodePackageProperties{
								{
									Name: to.Ptr("helloWorldCode"),
									Endpoints: []*armservicefabricmesh.EndpointProperties{
										{
											Name: to.Ptr("helloWorldListener"),
											Port: to.Ptr[int32](80),
										}},
									Image: to.Ptr("seabreeze/sbz-helloworld:1.0-alpine"),
									Resources: &armservicefabricmesh.ResourceRequirements{
										Requests: &armservicefabricmesh.ResourceRequests{
											CPU:        to.Ptr[float64](1),
											MemoryInGB: to.Ptr[float64](1),
										},
									},
								}},
							NetworkRefs: []*armservicefabricmesh.NetworkRef{
								{
									Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/networks/sampleNetwork"),
									EndpointRefs: []*armservicefabricmesh.EndpointRef{
										{
											Name: to.Ptr("helloWorldListener"),
										}},
								}},
							OSType: to.Ptr(armservicefabricmesh.OperatingSystemTypeLinux),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
