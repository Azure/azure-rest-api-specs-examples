package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97f789aeb52adfc1e20c386005839f5276874d7d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-05-02-preview/examples/ManagedNamespacesGet.json
func ExampleManagedNamespacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedNamespacesClient().Get(ctx, "rg1", "clustername1", "namespace1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedNamespace = armcontainerservice.ManagedNamespace{
	// 	Name: to.Ptr("namespace1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/managedNamespaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/managedNamespaces/namespace1"),
	// 	Location: to.Ptr("eastus2"),
	// 	Properties: &armcontainerservice.NamespaceProperties{
	// 		AdoptionPolicy: to.Ptr(armcontainerservice.AdoptionPolicyIfIdentical),
	// 		Annotations: map[string]*string{
	// 			"annatationKey": to.Ptr("annatationValue"),
	// 		},
	// 		DefaultNetworkPolicy: &armcontainerservice.NetworkPolicies{
	// 			Egress: to.Ptr(armcontainerservice.PolicyRuleAllowAll),
	// 			Ingress: to.Ptr(armcontainerservice.PolicyRuleAllowSameNamespace),
	// 		},
	// 		DefaultResourceQuota: &armcontainerservice.ResourceQuota{
	// 			CPULimit: to.Ptr("3m"),
	// 			CPURequest: to.Ptr("3m"),
	// 			MemoryLimit: to.Ptr("5Gi"),
	// 			MemoryRequest: to.Ptr("5Gi"),
	// 		},
	// 		DeletePolicy: to.Ptr(armcontainerservice.DeletePolicyKeep),
	// 		Labels: map[string]*string{
	// 			"kubernetes.azure.com/managedByArm": to.Ptr("true"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservice.NamespaceProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("tagValue1"),
	// 	},
	// }
}
