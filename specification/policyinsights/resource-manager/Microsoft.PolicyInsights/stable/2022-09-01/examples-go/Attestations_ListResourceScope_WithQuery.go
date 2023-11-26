package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-09-01/examples/Attestations_ListResourceScope_WithQuery.json
func ExampleAttestationsClient_NewListForResourcePager_listAttestationsAtIndividualResourceScopeWithQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAttestationsClient().NewListForResourcePager("subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5' AND PolicyDefinitionReferenceId eq '0b158b46-ff42-4799-8e39-08a5c23b4551'"),
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
		// page.AttestationListResult = armpolicyinsights.AttestationListResult{
		// 	Value: []*armpolicyinsights.Attestation{
		// 		{
		// 			Name: to.Ptr("790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
		// 			Type: to.Ptr("Microsoft.PolicyInsights/attestations"),
		// 			ID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM/providers/microsoft.policyinsights/attestations/790996e6-9871-4b1f-9cd9-ec42cd6ced1e"),
		// 			Properties: &armpolicyinsights.AttestationProperties{
		// 				AssessmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T00:00:00.000Z"); return t}()),
		// 				Comments: to.Ptr("This subscription has passed a security audit."),
		// 				ComplianceState: to.Ptr(armpolicyinsights.ComplianceStateCompliant),
		// 				Evidence: []*armpolicyinsights.AttestationEvidence{
		// 					{
		// 						Description: to.Ptr("The results of the security audit."),
		// 						SourceURI: to.Ptr("https://gist.github.com/contoso/9573e238762c60166c090ae16b814011"),
		// 				}},
		// 				ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T00:00:00.000Z"); return t}()),
		// 				LastComplianceStateChangeAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-15T18:52:27.000Z"); return t}()),
		// 				Metadata: map[string]any{
		// 					"departmentId": "NYC-MARKETING-1",
		// 				},
		// 				Owner: to.Ptr("55a32e28-3aa5-4eea-9b5a-4cd85153b966"),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
		// 				PolicyDefinitionReferenceID: to.Ptr("0b158b46-ff42-4799-8e39-08a5c23b4551"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SystemData: &armpolicyinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-15T18:52:27.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
		// 				CreatedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-15T18:52:27.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef6"),
		// 				LastModifiedByType: to.Ptr(armpolicyinsights.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
