package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_CreateSubscriptionScope.json
func ExampleAttestationsClient_BeginCreateOrUpdateAtSubscription_createAttestationAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAttestationsClient().BeginCreateOrUpdateAtSubscription(ctx, "790996e6-9871-4b1f-9cd9-ec42cd6ced1e", armpolicyinsights.Attestation{
		Properties: &armpolicyinsights.AttestationProperties{
			ComplianceState:    to.Ptr(armpolicyinsights.ComplianceStateCompliant),
			PolicyAssignmentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Attestation = armpolicyinsights.Attestation{
	// 	Name: to.Ptr("790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
	// 	Type: to.Ptr("Microsoft.PolicyInsights/attestations"),
	// 	ID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.policyinsights/attestations/790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
	// 	Properties: &armpolicyinsights.AttestationProperties{
	// 		ComplianceState: to.Ptr(armpolicyinsights.ComplianceStateCompliant),
	// 		LastComplianceStateChangeAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-15T18:52:27.000Z"); return t}()),
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// 	SystemData: &armpolicyinsights.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-15T18:52:27.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
	// 		CreatedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-15T18:52:27.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
	// 		LastModifiedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
	// 	},
	// }
}
