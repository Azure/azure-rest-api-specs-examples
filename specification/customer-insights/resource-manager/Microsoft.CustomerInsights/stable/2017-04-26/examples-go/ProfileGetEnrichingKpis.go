package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfileGetEnrichingKpis.json
func ExampleProfilesClient_GetEnrichingKpis() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().GetEnrichingKpis(ctx, "TestHubRG", "sdkTestHub", "TestProfileType396", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KpiDefinitionArray = []*armcustomerinsights.KpiDefinition{
	// 	{
	// 		Description: map[string]*string{
	// 			"en-us": to.Ptr("MonthlyDeposits"),
	// 		},
	// 		CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesMonth),
	// 		CalculationWindowFieldName: to.Ptr(""),
	// 		EntityType: to.Ptr(armcustomerinsights.EntityTypesInteraction),
	// 		EntityTypeName: to.Ptr("Deposit"),
	// 		Expression: to.Ptr("Amount"),
	// 		Filter: to.Ptr(""),
	// 		Function: to.Ptr(armcustomerinsights.KpiFunctionsSum),
	// 		GroupBy: []*string{
	// 			to.Ptr("AccountType"),
	// 			to.Ptr("BranchId"),
	// 			to.Ptr("ContactId"),
	// 			to.Ptr("Location"),
	// 			to.Ptr("Type")},
	// 			GroupByMetadata: []*armcustomerinsights.KpiGroupByMetadata{
	// 				{
	// 					DisplayName: map[string]*string{
	// 						"en-us": to.Ptr("AccountType"),
	// 					},
	// 					FieldName: to.Ptr("AccountType"),
	// 					FieldType: to.Ptr("Edm.String"),
	// 				},
	// 				{
	// 					DisplayName: map[string]*string{
	// 						"en-us": to.Ptr("BranchId"),
	// 					},
	// 					FieldName: to.Ptr("BranchId"),
	// 					FieldType: to.Ptr("Edm.Int32"),
	// 				},
	// 				{
	// 					DisplayName: map[string]*string{
	// 						"en-us": to.Ptr("ContactId"),
	// 					},
	// 					FieldName: to.Ptr("ContactId"),
	// 					FieldType: to.Ptr("Edm.Int32"),
	// 				},
	// 				{
	// 					DisplayName: map[string]*string{
	// 						"en-us": to.Ptr("Location"),
	// 					},
	// 					FieldName: to.Ptr("Location"),
	// 					FieldType: to.Ptr("Edm.String"),
	// 				},
	// 				{
	// 					DisplayName: map[string]*string{
	// 						"en-us": to.Ptr("Type"),
	// 					},
	// 					FieldName: to.Ptr("Type"),
	// 					FieldType: to.Ptr("Edm.String"),
	// 			}},
	// 			KpiName: to.Ptr("MonthlyDeposits"),
	// 			ParticipantProfilesMetadata: []*armcustomerinsights.KpiParticipantProfilesMetadata{
	// 			},
	// 			ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
	// 			TenantID: to.Ptr("sdkTestHub"),
	// 			ThresHolds: &armcustomerinsights.KpiThresholds{
	// 				IncreasingKpi: to.Ptr(true),
	// 				LowerLimit: to.Ptr[float64](0),
	// 				UpperLimit: to.Ptr[float64](0),
	// 			},
	// 			Unit: to.Ptr(""),
	// 	}}
}
