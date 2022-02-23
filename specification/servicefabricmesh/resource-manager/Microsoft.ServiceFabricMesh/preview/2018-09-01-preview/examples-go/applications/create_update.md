Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabricmesh%2Farmservicefabricmesh%2Fv0.2.1/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/create_update.json
func ExampleApplicationClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewApplicationClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<application-resource-name>",
		armservicefabricmesh.ApplicationResourceDescription{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.ApplicationResourceProperties{
				Description: to.StringPtr("<description>"),
				Services: []*armservicefabricmesh.ServiceResourceDescription{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armservicefabricmesh.ServiceResourceProperties{
							Description:  to.StringPtr("<description>"),
							ReplicaCount: to.Int32Ptr(1),
							CodePackages: []*armservicefabricmesh.ContainerCodePackageProperties{
								{
									Name: to.StringPtr("<name>"),
									Endpoints: []*armservicefabricmesh.EndpointProperties{
										{
											Name: to.StringPtr("<name>"),
											Port: to.Int32Ptr(80),
										}},
									Image: to.StringPtr("<image>"),
									Resources: &armservicefabricmesh.ResourceRequirements{
										Requests: &armservicefabricmesh.ResourceRequests{
											CPU:        to.Float64Ptr(1),
											MemoryInGB: to.Float64Ptr(1),
										},
									},
								}},
							NetworkRefs: []*armservicefabricmesh.NetworkRef{
								{
									Name: to.StringPtr("<name>"),
									EndpointRefs: []*armservicefabricmesh.EndpointRef{
										{
											Name: to.StringPtr("<name>"),
										}},
								}},
							OSType: armservicefabricmesh.OperatingSystemType("Linux").ToPtr(),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationClientCreateResult)
}
```
