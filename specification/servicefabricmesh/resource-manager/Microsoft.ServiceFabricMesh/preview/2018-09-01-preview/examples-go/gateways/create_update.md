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

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/create_update.json
func ExampleGatewayClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<gateway-resource-name>",
		armservicefabricmesh.GatewayResourceDescription{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.GatewayResourceProperties{
				Description: to.StringPtr("<description>"),
				DestinationNetwork: &armservicefabricmesh.NetworkRef{
					Name: to.StringPtr("<name>"),
				},
				SourceNetwork: &armservicefabricmesh.NetworkRef{
					Name: to.StringPtr("<name>"),
				},
				TCP: []*armservicefabricmesh.TCPConfig{
					{
						Name: to.StringPtr("<name>"),
						Destination: &armservicefabricmesh.GatewayDestination{
							ApplicationName: to.StringPtr("<application-name>"),
							EndpointName:    to.StringPtr("<endpoint-name>"),
							ServiceName:     to.StringPtr("<service-name>"),
						},
						Port: to.Int32Ptr(80),
					}},
				HTTP: []*armservicefabricmesh.HTTPConfig{
					{
						Name: to.StringPtr("<name>"),
						Hosts: []*armservicefabricmesh.HTTPHostConfig{
							{
								Name: to.StringPtr("<name>"),
								Routes: []*armservicefabricmesh.HTTPRouteConfig{
									{
										Name: to.StringPtr("<name>"),
										Destination: &armservicefabricmesh.GatewayDestination{
											ApplicationName: to.StringPtr("<application-name>"),
											EndpointName:    to.StringPtr("<endpoint-name>"),
											ServiceName:     to.StringPtr("<service-name>"),
										},
										Match: &armservicefabricmesh.HTTPRouteMatchRule{
											Path: &armservicefabricmesh.HTTPRouteMatchPath{
												Type:    armservicefabricmesh.PathMatchType("prefix").ToPtr(),
												Rewrite: to.StringPtr("<rewrite>"),
												Value:   to.StringPtr("<value>"),
											},
											Headers: []*armservicefabricmesh.HTTPRouteMatchHeader{
												{
													Name:  to.StringPtr("<name>"),
													Type:  armservicefabricmesh.HeaderMatchType("exact").ToPtr(),
													Value: to.StringPtr("<value>"),
												}},
										},
									}},
							}},
						Port: to.Int32Ptr(8081),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GatewayClientCreateResult)
}
```
