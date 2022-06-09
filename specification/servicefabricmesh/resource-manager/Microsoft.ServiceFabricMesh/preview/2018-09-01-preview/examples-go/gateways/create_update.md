```go
package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/create_update.json
func ExampleGatewayClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabricmesh.NewGatewayClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"sbz_demo",
		"sampleGateway",
		armservicefabricmesh.GatewayResourceDescription{
			Location: to.Ptr("EastUS"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.GatewayResourceProperties{
				Description: to.Ptr("Service Fabric Mesh sample gateway."),
				DestinationNetwork: &armservicefabricmesh.NetworkRef{
					Name: to.Ptr("helloWorldNetwork"),
				},
				SourceNetwork: &armservicefabricmesh.NetworkRef{
					Name: to.Ptr("Open"),
				},
				TCP: []*armservicefabricmesh.TCPConfig{
					{
						Name: to.Ptr("web"),
						Destination: &armservicefabricmesh.GatewayDestination{
							ApplicationName: to.Ptr("helloWorldApp"),
							EndpointName:    to.Ptr("helloWorldListener"),
							ServiceName:     to.Ptr("helloWorldService"),
						},
						Port: to.Ptr[int32](80),
					}},
				HTTP: []*armservicefabricmesh.HTTPConfig{
					{
						Name: to.Ptr("contosoWebsite"),
						Hosts: []*armservicefabricmesh.HTTPHostConfig{
							{
								Name: to.Ptr("contoso.com"),
								Routes: []*armservicefabricmesh.HTTPRouteConfig{
									{
										Name: to.Ptr("index"),
										Destination: &armservicefabricmesh.GatewayDestination{
											ApplicationName: to.Ptr("httpHelloWorldApp"),
											EndpointName:    to.Ptr("indexHttpEndpoint"),
											ServiceName:     to.Ptr("indexService"),
										},
										Match: &armservicefabricmesh.HTTPRouteMatchRule{
											Path: &armservicefabricmesh.HTTPRouteMatchPath{
												Type:    to.Ptr(armservicefabricmesh.PathMatchTypePrefix),
												Rewrite: to.Ptr("/"),
												Value:   to.Ptr("/index"),
											},
											Headers: []*armservicefabricmesh.HTTPRouteMatchHeader{
												{
													Name:  to.Ptr("accept"),
													Type:  to.Ptr(armservicefabricmesh.HeaderMatchTypeExact),
													Value: to.Ptr("application/json"),
												}},
										},
									}},
							}},
						Port: to.Ptr[int32](8081),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabricmesh%2Farmservicefabricmesh%2Fv0.5.0/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh/README.md) on how to add the SDK to your project and authenticate.
