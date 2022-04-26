Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabricmesh%2Farmservicefabricmesh%2Fv0.4.0/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh/README.md) on how to add the SDK to your project and authenticate.

```go
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
		return
	}
	ctx := context.Background()
	client, err := armservicefabricmesh.NewApplicationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<application-resource-name>",
		armservicefabricmesh.ApplicationResourceDescription{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.ApplicationResourceProperties{
				Description: to.Ptr("<description>"),
				Services: []*armservicefabricmesh.ServiceResourceDescription{
					{
						Name: to.Ptr("<name>"),
						Properties: &armservicefabricmesh.ServiceResourceProperties{
							Description:  to.Ptr("<description>"),
							ReplicaCount: to.Ptr[int32](1),
							CodePackages: []*armservicefabricmesh.ContainerCodePackageProperties{
								{
									Name: to.Ptr("<name>"),
									Endpoints: []*armservicefabricmesh.EndpointProperties{
										{
											Name: to.Ptr("<name>"),
											Port: to.Ptr[int32](80),
										}},
									Image: to.Ptr("<image>"),
									Resources: &armservicefabricmesh.ResourceRequirements{
										Requests: &armservicefabricmesh.ResourceRequests{
											CPU:        to.Ptr[float64](1),
											MemoryInGB: to.Ptr[float64](1),
										},
									},
								}},
							NetworkRefs: []*armservicefabricmesh.NetworkRef{
								{
									Name: to.Ptr("<name>"),
									EndpointRefs: []*armservicefabricmesh.EndpointRef{
										{
											Name: to.Ptr("<name>"),
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
		return
	}
	// TODO: use response item
	_ = res
}
```
