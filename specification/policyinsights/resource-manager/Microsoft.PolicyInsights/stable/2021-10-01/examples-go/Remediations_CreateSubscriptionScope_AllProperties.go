package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CreateSubscriptionScope_AllProperties.json
func ExampleRemediationsClient_CreateOrUpdateAtSubscription_createRemediationAtSubscriptionScopeWithAllProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewRemediationsClient("35ee058e-5fa0-414c-8145-3ebb8d09b6e2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateAtSubscription(ctx, "storageRemediation", armpolicyinsights.Remediation{
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
	// TODO: use response item
	_ = res
}
