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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/create_update.json
func ExampleGatewayClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<gateway-resource-name>",
		armservicefabricmesh.GatewayResourceDescription{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabricmesh.GatewayResourceProperties{
				Description: to.Ptr("<description>"),
				DestinationNetwork: &armservicefabricmesh.NetworkRef{
					Name: to.Ptr("<name>"),
				},
				SourceNetwork: &armservicefabricmesh.NetworkRef{
					Name: to.Ptr("<name>"),
				},
				TCP: []*armservicefabricmesh.TCPConfig{
					{
						Name: to.Ptr("<name>"),
						Destination: &armservicefabricmesh.GatewayDestination{
							ApplicationName: to.Ptr("<application-name>"),
							EndpointName:    to.Ptr("<endpoint-name>"),
							ServiceName:     to.Ptr("<service-name>"),
						},
						Port: to.Ptr[int32](80),
					}},
				HTTP: []*armservicefabricmesh.HTTPConfig{
					{
						Name: to.Ptr("<name>"),
						Hosts: []*armservicefabricmesh.HTTPHostConfig{
							{
								Name: to.Ptr("<name>"),
								Routes: []*armservicefabricmesh.HTTPRouteConfig{
									{
										Name: to.Ptr("<name>"),
										Destination: &armservicefabricmesh.GatewayDestination{
											ApplicationName: to.Ptr("<application-name>"),
											EndpointName:    to.Ptr("<endpoint-name>"),
											ServiceName:     to.Ptr("<service-name>"),
										},
										Match: &armservicefabricmesh.HTTPRouteMatchRule{
											Path: &armservicefabricmesh.HTTPRouteMatchPath{
												Type:    to.Ptr(armservicefabricmesh.PathMatchTypePrefix),
												Rewrite: to.Ptr("<rewrite>"),
												Value:   to.Ptr("<value>"),
											},
											Headers: []*armservicefabricmesh.HTTPRouteMatchHeader{
												{
													Name:  to.Ptr("<name>"),
													Type:  to.Ptr(armservicefabricmesh.HeaderMatchTypeExact),
													Value: to.Ptr("<value>"),
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
		return
	}
	// TODO: use response item
	_ = res
}
```
