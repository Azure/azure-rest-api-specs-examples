package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2025-08-01-preview/FleetManagedNamespaces_ListByFleet.json
func ExampleFleetManagedNamespacesClient_NewListByFleetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetManagedNamespacesClient().NewListByFleetPager("rgfleets", "fleet1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerservicefleet.FleetManagedNamespacesClientListByFleetResponse{
		// 	FleetManagedNamespaceListResult: armcontainerservicefleet.FleetManagedNamespaceListResult{
		// 		Value: []*armcontainerservicefleet.FleetManagedNamespace{
		// 			{
		// 				Properties: &armcontainerservicefleet.FleetManagedNamespaceProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.FleetManagedNamespaceProvisioningStateSucceeded),
		// 					ManagedNamespaceProperties: &armcontainerservicefleet.ManagedNamespaceProperties{
		// 						Labels: map[string]*string{
		// 							"key1": to.Ptr("value1"),
		// 						},
		// 						Annotations: map[string]*string{
		// 							"key2": to.Ptr("value2"),
		// 						},
		// 						DefaultResourceQuota: &armcontainerservicefleet.ResourceQuota{
		// 							CPURequest: to.Ptr("1"),
		// 							CPULimit: to.Ptr("1"),
		// 							MemoryRequest: to.Ptr("10Gi"),
		// 							MemoryLimit: to.Ptr("32Gi"),
		// 						},
		// 						DefaultNetworkPolicy: &armcontainerservicefleet.NetworkPolicy{
		// 							Ingress: to.Ptr(armcontainerservicefleet.PolicyRuleAllowSameNamespace),
		// 							Egress: to.Ptr(armcontainerservicefleet.PolicyRuleAllowAll),
		// 						},
		// 					},
		// 					AdoptionPolicy: to.Ptr(armcontainerservicefleet.AdoptionPolicyNever),
		// 					DeletePolicy: to.Ptr(armcontainerservicefleet.DeletePolicyKeep),
		// 					PropagationPolicy: &armcontainerservicefleet.PropagationPolicy{
		// 						Type: to.Ptr(armcontainerservicefleet.PropagationTypePlacement),
		// 						PlacementProfile: &armcontainerservicefleet.PlacementProfile{
		// 							DefaultClusterResourcePlacement: &armcontainerservicefleet.ClusterResourcePlacementSpec{
		// 								Policy: &armcontainerservicefleet.PlacementPolicy{
		// 									PlacementType: to.Ptr(armcontainerservicefleet.PlacementTypePickAll),
		// 									Affinity: &armcontainerservicefleet.Affinity{
		// 										ClusterAffinity: &armcontainerservicefleet.ClusterAffinity{
		// 											RequiredDuringSchedulingIgnoredDuringExecution: &armcontainerservicefleet.ClusterSelector{
		// 												ClusterSelectorTerms: []*armcontainerservicefleet.ClusterSelectorTerm{
		// 													{
		// 														LabelSelector: &armcontainerservicefleet.LabelSelector{
		// 															MatchLabels: map[string]*string{
		// 																"gpu": to.Ptr("true"),
		// 															},
		// 															MatchExpressions: []*armcontainerservicefleet.LabelSelectorRequirement{
		// 																{
		// 																	Key: to.Ptr("region"),
		// 																	Operator: to.Ptr(armcontainerservicefleet.LabelSelectorOperatorIn),
		// 																	Values: []*string{
		// 																		to.Ptr("production1"),
		// 																		to.Ptr("production2"),
		// 																	},
		// 																},
		// 															},
		// 														},
		// 														PropertySelector: &armcontainerservicefleet.PropertySelector{
		// 															MatchExpressions: []*armcontainerservicefleet.PropertySelectorRequirement{
		// 																{
		// 																	Name: to.Ptr("zones"),
		// 																	Operator: to.Ptr(armcontainerservicefleet.PropertySelectorOperatorGt),
		// 																	Values: []*string{
		// 																		to.Ptr("1"),
		// 																	},
		// 																},
		// 															},
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 									Tolerations: []*armcontainerservicefleet.Toleration{
		// 										{
		// 											Key: to.Ptr("AIWorkloadOnly"),
		// 											Operator: to.Ptr(armcontainerservicefleet.TolerationOperatorExists),
		// 											Value: to.Ptr("true"),
		// 											Effect: to.Ptr(armcontainerservicefleet.TaintEffectNoSchedule),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Status: &armcontainerservicefleet.FleetManagedNamespaceStatus{
		// 						LastOperationError: &armcontainerservicefleet.ErrorDetail{
		// 						},
		// 					},
		// 					PortalFqdn: to.Ptr("fleet1-namespace1-abc123.flt.eastus.azmk8s.io"),
		// 				},
		// 				ETag: to.Ptr("\"EtagValue\""),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("tagValue1"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet1/managedNamespaces/namespace1"),
		// 				Name: to.Ptr("namespace1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets/managedNamespaces"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("someUser"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-23T05:40:40.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("someOtherUser"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-23T05:40:40.657Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://nextlink.sausage.com"),
		// 	},
		// }
	}
}
