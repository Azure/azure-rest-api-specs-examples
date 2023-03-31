package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/get.json
func ExampleApplicationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmesh.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationClient().Get(ctx, "sbz_demo", "sampleApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationResourceDescription = armservicefabricmesh.ApplicationResourceDescription{
	// 	Name: to.Ptr("sampleApplication"),
	// 	Type: to.Ptr("Microsoft.ServiceFabricMesh/applications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/applications/sampleApplication"),
	// 	Location: to.Ptr("EastUS"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabricmesh.ApplicationResourceProperties{
	// 		Description: to.Ptr("Service Fabric Mesh sample application."),
	// 		HealthState: to.Ptr(armservicefabricmesh.HealthStateOk),
	// 		ServiceNames: []*string{
	// 			to.Ptr("helloWorldService")},
	// 			Status: to.Ptr(armservicefabricmesh.ResourceStatusReady),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}
