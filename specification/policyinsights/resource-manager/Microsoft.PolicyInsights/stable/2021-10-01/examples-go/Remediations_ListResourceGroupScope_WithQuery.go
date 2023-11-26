package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListResourceGroupScope_WithQuery.json
func ExampleRemediationsClient_NewListForResourceGroupPager_listRemediationsAtResourceGroupScopeWithQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRemediationsClient().NewListForResourceGroupPager("myResourceGroup", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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
		// page.RemediationListResult = armpolicyinsights.RemediationListResult{
		// 	Value: []*armpolicyinsights.Remediation{
		// 		{
		// 			Name: to.Ptr("790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
		// 			Type: to.Ptr("Microsoft.PolicyInsights/remediations"),
		// 			ID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.policyinsights/remediations/790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
		// 			Properties: &armpolicyinsights.RemediationProperties{
		// 				CorrelationID: to.Ptr("a14e1d60-dae9-4771-b4be-a556d69e77a6"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:51:09.075Z"); return t}()),
		// 				DeploymentStatus: &armpolicyinsights.RemediationDeploymentSummary{
		// 					FailedDeployments: to.Ptr[int32](0),
		// 					SuccessfulDeployments: to.Ptr[int32](2),
		// 					TotalDeployments: to.Ptr[int32](2),
		// 				},
		// 				FailureThreshold: &armpolicyinsights.RemediationPropertiesFailureThreshold{
		// 					Percentage: to.Ptr[float32](0.1),
		// 				},
		// 				LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:52:10.001Z"); return t}()),
		// 				ParallelDeployments: to.Ptr[int32](6),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceCount: to.Ptr[int32](42),
		// 				ResourceDiscoveryMode: to.Ptr(armpolicyinsights.ResourceDiscoveryModeExistingNonCompliant),
		// 				StatusMessage: to.Ptr("Remediation extended status"),
		// 			},
		// 			SystemData: &armpolicyinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:49:16.463Z"); return t}()),
		// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
		// 				CreatedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:50:18.464Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
		// 				LastModifiedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("0866f0cc-6cb3-422c-9fb6-a4e3e2266da8"),
		// 			Type: to.Ptr("Microsoft.PolicyInsights/remediations"),
		// 			ID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.policyinsights/remediations/0866f0cc-6cb3-422c-9fb6-a4e3e2266da8"),
		// 			Properties: &armpolicyinsights.RemediationProperties{
		// 				CorrelationID: to.Ptr("957378d3-a143-421e-8b0e-3048dff01932"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:49:16.463Z"); return t}()),
		// 				DeploymentStatus: &armpolicyinsights.RemediationDeploymentSummary{
		// 					FailedDeployments: to.Ptr[int32](0),
		// 					SuccessfulDeployments: to.Ptr[int32](2),
		// 					TotalDeployments: to.Ptr[int32](2),
		// 				},
		// 				LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:50:18.464Z"); return t}()),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceDiscoveryMode: to.Ptr(armpolicyinsights.ResourceDiscoveryModeExistingNonCompliant),
		// 				StatusMessage: to.Ptr("Remediation extended status"),
		// 			},
		// 			SystemData: &armpolicyinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:49:16.463Z"); return t}()),
		// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
		// 				CreatedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:50:18.464Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
		// 				LastModifiedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
