package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/replicas/get.json
func ExampleServiceReplicaClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmesh.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceReplicaClient().Get(ctx, "sbz_demo", "helloWorldApp", "helloWorldService", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceReplicaDescription = armservicefabricmesh.ServiceReplicaDescription{
	// 	CodePackages: []*armservicefabricmesh.ContainerCodePackageProperties{
	// 		{
	// 			Name: to.Ptr("helloWorldCode"),
	// 			Endpoints: []*armservicefabricmesh.EndpointProperties{
	// 				{
	// 					Name: to.Ptr("helloWorldListener"),
	// 					Port: to.Ptr[int32](80),
	// 			}},
	// 			Image: to.Ptr("seabreeze/sbz-helloworld:1.0-alpine"),
	// 			InstanceView: &armservicefabricmesh.ContainerInstanceView{
	// 				CurrentState: &armservicefabricmesh.ContainerState{
	// 					ExitCode: to.Ptr("0"),
	// 					State: to.Ptr("Running"),
	// 				},
	// 				Events: []*armservicefabricmesh.ContainerEvent{
	// 					{
	// 						Name: to.Ptr("Created"),
	// 						Type: to.Ptr("Normal"),
	// 						Count: to.Ptr[int32](3),
	// 						FirstTimestamp: to.Ptr("2018-04-05T22:37:20.9016844"),
	// 						LastTimestamp: to.Ptr("2018-04-06T06:36:06.0887046"),
	// 						Message: to.Ptr("Container created and started."),
	// 					},
	// 					{
	// 						Name: to.Ptr("Stopped"),
	// 						Type: to.Ptr("Normal"),
	// 						Count: to.Ptr[int32](1),
	// 						FirstTimestamp: to.Ptr("2018-04-06T06:34:00.6622454"),
	// 						LastTimestamp: to.Ptr("2018-04-06T06:34:00.6622454"),
	// 						Message: to.Ptr("Container was stopped."),
	// 				}},
	// 				PreviousState: &armservicefabricmesh.ContainerState{
	// 					ExitCode: to.Ptr("0"),
	// 					State: to.Ptr("NotSpecified"),
	// 				},
	// 				RestartCount: to.Ptr[int32](1),
	// 			},
	// 			Resources: &armservicefabricmesh.ResourceRequirements{
	// 				Requests: &armservicefabricmesh.ResourceRequests{
	// 					CPU: to.Ptr[float64](1),
	// 					MemoryInGB: to.Ptr[float64](1),
	// 				},
	// 			},
	// 	}},
	// 	NetworkRefs: []*armservicefabricmesh.NetworkRef{
	// 		{
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/networks/sampleNetwork"),
	// 			EndpointRefs: []*armservicefabricmesh.EndpointRef{
	// 				{
	// 					Name: to.Ptr("helloWorldListener"),
	// 			}},
	// 	}},
	// 	OSType: to.Ptr(armservicefabricmesh.OperatingSystemTypeLinux),
	// 	ReplicaName: to.Ptr("0"),
	// }
}
