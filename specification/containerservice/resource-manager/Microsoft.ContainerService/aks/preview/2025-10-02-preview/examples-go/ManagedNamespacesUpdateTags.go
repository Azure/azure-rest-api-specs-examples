package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/ManagedNamespacesUpdateTags.json
func ExampleManagedNamespacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedNamespacesClient().Update(ctx, "rg1", "clustername1", "namespace1", armcontainerservice.TagsObject{
		Tags: map[string]*string{
			"tagKey1": to.Ptr("tagValue1"),
			"tagKey2": to.Ptr("tagValue2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedNamespace = armcontainerservice.ManagedNamespace{
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
	// 		"tagKey2": to.Ptr("tagValue2"),
	// 	},
	// }
}
