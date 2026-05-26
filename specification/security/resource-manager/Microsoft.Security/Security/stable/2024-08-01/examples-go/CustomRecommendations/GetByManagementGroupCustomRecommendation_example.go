package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2024-08-01/CustomRecommendations/GetByManagementGroupCustomRecommendation_example.json
func ExampleCustomRecommendationsClient_Get_getACustomRecommendationOverManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomRecommendationsClient().Get(ctx, "providers/Microsoft.Management/managementGroups/contoso", "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.CustomRecommendationsClientGetResponse{
	// 	CustomRecommendation: armsecurity.CustomRecommendation{
	// 		Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Type: to.Ptr("Microsoft.Security/customRecommendations"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/contoso/providers/Microsoft.Security/customRecommendations/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Properties: &armsecurity.CustomRecommendationProperties{
	// 			Description: to.Ptr("organization passwords policy"),
	// 			AssessmentKey: to.Ptr("d5f442f7-7e77-4bcf-a450-a9c1b9a94eeb"),
	// 			CloudProviders: []*armsecurity.RecommendationSupportedClouds{
	// 				to.Ptr(armsecurity.RecommendationSupportedCloudsAWS),
	// 			},
	// 			DisplayName: to.Ptr("Password Policy"),
	// 			Query: to.Ptr("RawEntityMetadata | where Environment == 'GCP' and Identifiers.Type == 'compute.firewalls' | extend IslogConfigEnabled = tobool(Record.logConfig.enable) | extend HealthStatus = iff(IslogConfigEnabled, 'HEALTHY', 'UNHEALTHY')"),
	// 			RemediationDescription: to.Ptr("Change password policy to..."),
	// 			SecurityIssue: to.Ptr(armsecurity.SecurityIssueVulnerability),
	// 			Severity: to.Ptr(armsecurity.SeverityEnumMedium),
	// 		},
	// 		SystemData: &armsecurity.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
