package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/get.json
func ExampleServiceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmesh.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().Get(ctx, "sbz_demo", "sampleApplication", "helloWorldService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResourceDescription = armservicefabricmesh.ServiceResourceDescription{
	// 	Name: to.Ptr("helloWorldService"),
	// 	Type: to.Ptr("Microsoft.ServiceFabricMesh/services"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/applications/sampleApplication/services/helloWorldService"),
	// 	Properties: &armservicefabricmesh.ServiceResourceProperties{
	// 		Description: to.Ptr("SeaBreeze Hello World Service."),
	// 		HealthState: to.Ptr(armservicefabricmesh.HealthStateOk),
	// 		ReplicaCount: to.Ptr[int32](1),
	// 		Status: to.Ptr(armservicefabricmesh.ResourceStatusReady),
	// 		CodePackages: []*armservicefabricmesh.ContainerCodePackageProperties{
	// 			{
	// 				Name: to.Ptr("helloWorldCode"),
	// 				Endpoints: []*armservicefabricmesh.EndpointProperties{
	// 					{
	// 						Name: to.Ptr("helloWorldListener"),
	// 						Port: to.Ptr[int32](80),
	// 				}},
	// 				Image: to.Ptr("seabreeze/sbz-helloworld:1.0-alpine"),
	// 				Resources: &armservicefabricmesh.ResourceRequirements{
	// 					Requests: &armservicefabricmesh.ResourceRequests{
	// 						CPU: to.Ptr[float64](1),
	// 						MemoryInGB: to.Ptr[float64](1),
	// 					},
	// 				},
	// 		}},
	// 		NetworkRefs: []*armservicefabricmesh.NetworkRef{
	// 			{
	// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/networks/sampleNetwork"),
	// 				EndpointRefs: []*armservicefabricmesh.EndpointRef{
	// 					{
	// 						Name: to.Ptr("helloWorldListener"),
	// 				}},
	// 		}},
	// 		OSType: to.Ptr(armservicefabricmesh.OperatingSystemTypeLinux),
	// 	},
	// }
}
