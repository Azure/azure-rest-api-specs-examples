package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/get.json
func ExampleGatewayClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmesh.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().Get(ctx, "sbz_demo", "sampleGateway", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayResourceDescription = armservicefabricmesh.GatewayResourceDescription{
	// 	Name: to.Ptr("sampleGateway"),
	// 	Type: to.Ptr("Microsoft.ServiceFabricMesh/gateways"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/gateways/sampleGateway"),
	// 	Location: to.Ptr("EastUS"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabricmesh.GatewayResourceProperties{
	// 		Description: to.Ptr("Service Fabric Mesh sample gateway."),
	// 		DestinationNetwork: &armservicefabricmesh.NetworkRef{
	// 			Name: to.Ptr("helloWorldNetwork"),
	// 		},
	// 		IPAddress: to.Ptr("192.168.1.1"),
	// 		SourceNetwork: &armservicefabricmesh.NetworkRef{
	// 			Name: to.Ptr("Open"),
	// 		},
	// 		Status: to.Ptr(armservicefabricmesh.ResourceStatusReady),
	// 		TCP: []*armservicefabricmesh.TCPConfig{
	// 			{
	// 				Name: to.Ptr("web"),
	// 				Destination: &armservicefabricmesh.GatewayDestination{
	// 					ApplicationName: to.Ptr("helloWorldApp"),
	// 					EndpointName: to.Ptr("helloWorldListener"),
	// 					ServiceName: to.Ptr("helloWorldService"),
	// 				},
	// 				Port: to.Ptr[int32](80),
	// 		}},
	// 		HTTP: []*armservicefabricmesh.HTTPConfig{
	// 			{
	// 				Name: to.Ptr("contosoWebsite"),
	// 				Hosts: []*armservicefabricmesh.HTTPHostConfig{
	// 					{
	// 						Name: to.Ptr("contoso.com"),
	// 						Routes: []*armservicefabricmesh.HTTPRouteConfig{
	// 							{
	// 								Name: to.Ptr("index"),
	// 								Destination: &armservicefabricmesh.GatewayDestination{
	// 									ApplicationName: to.Ptr("httpHelloWorldApp"),
	// 									EndpointName: to.Ptr("indexHttpEndpoint"),
	// 									ServiceName: to.Ptr("indexService"),
	// 								},
	// 								Match: &armservicefabricmesh.HTTPRouteMatchRule{
	// 									Path: &armservicefabricmesh.HTTPRouteMatchPath{
	// 										Type: to.Ptr(armservicefabricmesh.PathMatchTypePrefix),
	// 										Rewrite: to.Ptr("/"),
	// 										Value: to.Ptr("/index"),
	// 									},
	// 									Headers: []*armservicefabricmesh.HTTPRouteMatchHeader{
	// 										{
	// 											Name: to.Ptr("accept"),
	// 											Type: to.Ptr(armservicefabricmesh.HeaderMatchTypeExact),
	// 											Value: to.Ptr("application/json"),
	// 									}},
	// 								},
	// 						}},
	// 				}},
	// 				Port: to.Ptr[int32](8081),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
