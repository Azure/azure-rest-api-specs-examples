package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/ListSubAssessments_example.json
func ExampleSubAssessmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubAssessmentsClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "82e20e14-edc5-4373-bfc4-f13121257c37", nil)
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
		// page.SubAssessmentList = armsecurity.SubAssessmentList{
		// 	Value: []*armsecurity.SubAssessment{
		// 		{
		// 			Name: to.Ptr("8fbe5054-e97c-3a7a-fda7-c8308ca8d3cf"),
		// 			Type: to.Ptr("Microsoft.Security/assessments/subAssessments"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/ascdemoRG/providers/Microsoft.Sql/servers/sqlserver1demo/providers/Microsoft.Security/assessments/82e20e14-edc5-4373-bfc4-f13121257c37/subassessments/8fbe5054-e97c-3a7a-fda7-c8308ca8d3cf"),
		// 			Properties: &armsecurity.SubAssessmentProperties{
		// 				Description: to.Ptr("The Azure SQL Database-level firewall helps protect your data by preventing all access to your database until you specify which IP addresses have permission. Database-level firewall rules grant access to the specific database based on the originating IP address of each request.\n\nDatabase-level firewall rules for master"),
		// 				AdditionalData: &armsecurity.SQLServerVulnerabilityProperties{
		// 					AssessedResourceType: to.Ptr(armsecurity.AssessedResourceTypeSQLServerVulnerability),
		// 					Type: to.Ptr("AzureDatabase"),
		// 					Query: to.Ptr("SELECT name\n    ,start_ip_address\n    ,end_ip_address\nFROM sys.database_firewall_rules"),
		// 				},
		// 				Category: to.Ptr("SurfaceAreaReduction"),
		// 				DisplayName: to.Ptr("Database-level firewall rules should be tracked and maintained at a strict minimum"),
		// 				ID: to.Ptr("VA2064"),
		// 				Impact: to.Ptr("Firewall rules should be strictly configured to allow access only to client computers that have a valid need to connect to the database. Any superfluous entries in the firewall may pose a threat by allowing an unauthorized source access to your database."),
		// 				Remediation: to.Ptr("Evaluate each of the database-level firewall rules. Remove any rules that grant unnecessary access and set the rest as a baseline. Deviations from the baseline will be identified and brought to your attention in subsequent scans."),
		// 				ResourceDetails: &armsecurity.AzureResourceDetails{
		// 					Source: to.Ptr(armsecurity.SourceAzure),
		// 					ID: to.Ptr("/subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/ascdemoRG/providers/Microsoft.Sql/servers/sqlserver1demo/databases/database1"),
		// 				},
		// 				Status: &armsecurity.SubAssessmentStatus{
		// 					Cause: to.Ptr("Unknown"),
		// 					Code: to.Ptr(armsecurity.SubAssessmentStatusCodeHealthy),
		// 					Severity: to.Ptr(armsecurity.SeverityHigh),
		// 				},
		// 				TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-23T12:20:08.764Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
