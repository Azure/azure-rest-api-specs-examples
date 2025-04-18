package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b7b354f5238cfab8b22a1bfa992a12ecca42ca1b/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-01-02-preview/examples/LoadBalancers_Create_Or_Update.json
func ExampleLoadBalancersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancersClient().CreateOrUpdate(ctx, "rg1", "clustername1", "kubernetes", armcontainerservice.LoadBalancer{
		Properties: &armcontainerservice.LoadBalancerProperties{
			AllowServicePlacement: to.Ptr(true),
			PrimaryAgentPoolName:  to.Ptr("agentpool1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoadBalancer = armcontainerservice.LoadBalancer{
	// 	Name: to.Ptr("kubernetes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/loadBalancers/kubernetes"),
	// 	Properties: &armcontainerservice.LoadBalancerProperties{
	// 		AllowServicePlacement: to.Ptr(true),
	// 		PrimaryAgentPoolName: to.Ptr("agentPool1"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
