package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CreateSubscriptionScope_AllProperties.json
func ExampleRemediationsClient_CreateOrUpdateAtSubscription_createRemediationAtSubscriptionScopeWithAllProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRemediationsClient().CreateOrUpdateAtSubscription(ctx, "storageRemediation", armpolicyinsights.Remediation{
		Properties: &armpolicyinsights.RemediationProperties{
			FailureThreshold: &armpolicyinsights.RemediationPropertiesFailureThreshold{
				Percentage: to.Ptr[float32](0.1),
			},
			Filters: &armpolicyinsights.RemediationFilters{
				Locations: []*string{
					to.Ptr("eastus"),
					to.Ptr("westus")},
			},
			ParallelDeployments:         to.Ptr[int32](6),
			PolicyAssignmentID:          to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
			PolicyDefinitionReferenceID: to.Ptr("8c8fa9e4"),
			ResourceCount:               to.Ptr[int32](42),
			ResourceDiscoveryMode:       to.Ptr(armpolicyinsights.ResourceDiscoveryModeReEvaluateCompliance),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Remediation = armpolicyinsights.Remediation{
	// 	Name: to.Ptr("storageRemediation"),
	// 	Type: to.Ptr("Microsoft.PolicyInsights/remediations"),
	// 	ID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.policyinsights/remediations/storageRemediation"),
	// 	Properties: &armpolicyinsights.RemediationProperties{
	// 		CorrelationID: to.Ptr("a14e1d60-dae9-4771-b4be-a556d69e77a6"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-12T20:58:58.753Z"); return t}()),
	// 		DeploymentStatus: &armpolicyinsights.RemediationDeploymentSummary{
	// 			FailedDeployments: to.Ptr[int32](0),
	// 			SuccessfulDeployments: to.Ptr[int32](0),
	// 			TotalDeployments: to.Ptr[int32](2),
	// 		},
	// 		FailureThreshold: &armpolicyinsights.RemediationPropertiesFailureThreshold{
	// 			Percentage: to.Ptr[float32](0.1),
	// 		},
	// 		LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-12T20:58:58.753Z"); return t}()),
	// 		ParallelDeployments: to.Ptr[int32](6),
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
	// 		PolicyDefinitionReferenceID: to.Ptr("8c8fa9e4"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceCount: to.Ptr[int32](42),
	// 		ResourceDiscoveryMode: to.Ptr(armpolicyinsights.ResourceDiscoveryModeReEvaluateCompliance),
	// 		StatusMessage: to.Ptr("Remediation extended status"),
	// 	},
	// 	SystemData: &armpolicyinsights.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-12T20:58:58.753Z"); return t}()),
	// 		CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
	// 		CreatedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-12T20:58:58.753Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
	// 		LastModifiedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
	// 	},
	// }
}
