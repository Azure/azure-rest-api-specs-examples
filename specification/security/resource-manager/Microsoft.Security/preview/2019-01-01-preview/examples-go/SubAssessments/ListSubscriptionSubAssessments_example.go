package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/ListSubscriptionSubAssessments_example.json
func ExampleSubAssessmentsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubAssessmentsClient().NewListAllPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", nil)
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
		// 			Name: to.Ptr("8c98f353-8b41-4e77-979b-6adeecd5d168"),
		// 			Type: to.Ptr("Microsoft.Security/assessments/subAssessments"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.ContainerRegistry/registries/myRegistry/providers/Microsoft.Security/assessments/dbd0cb49-b563-45e7-9724-889e799fa648/subAssessments/8c98f353-8b41-4e77-979b-6adeecd5d168"),
		// 			Properties: &armsecurity.SubAssessmentProperties{
		// 				Description: to.Ptr("The backdoor 'Back Orifice' was detected on this system.  The presence of this backdoor indicates that your system has already been compromised.  Unauthorized users can access your host at any time. Unauthorized users can take complete control of the host and manipulate data.  They can steal the data or even wipe out the host."),
		// 				AdditionalData: &armsecurity.ContainerRegistryVulnerabilityProperties{
		// 					AssessedResourceType: to.Ptr(armsecurity.AssessedResourceTypeContainerRegistryVulnerability),
		// 					Type: to.Ptr("Vulnerability"),
		// 					Cve: []*armsecurity.CVE{
		// 						{
		// 							Link: to.Ptr("http://contoso.com"),
		// 							Title: to.Ptr("CVE-2019-12345"),
		// 					}},
		// 					Cvss: map[string]*armsecurity.CVSS{
		// 						"2.0": &armsecurity.CVSS{
		// 							Base: to.Ptr[float32](10),
		// 						},
		// 						"3.0": &armsecurity.CVSS{
		// 							Base: to.Ptr[float32](10),
		// 						},
		// 					},
		// 					ImageDigest: to.Ptr("c186fc44-3154-4ce2-ba18-b719d895c3b0"),
		// 					Patchable: to.Ptr(true),
		// 					PublishedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.0000000Z"); return t}()),
		// 					RepositoryName: to.Ptr("myRepo"),
		// 					VendorReferences: []*armsecurity.VendorReference{
		// 						{
		// 							Link: to.Ptr("http://contoso.com"),
		// 							Title: to.Ptr("Reference_1"),
		// 					}},
		// 				},
		// 				Category: to.Ptr("Backdoors and trojan horses"),
		// 				DisplayName: to.Ptr("'Back Orifice' Backdoor"),
		// 				ID: to.Ptr("1001"),
		// 				Impact: to.Ptr("3"),
		// 				Remediation: to.Ptr("Use a recent anti-virus program to remove this backdoor and check your system regularly with anti-virus software."),
		// 				ResourceDetails: &armsecurity.AzureResourceDetails{
		// 					Source: to.Ptr(armsecurity.SourceAzure),
		// 					ID: to.Ptr("repositories/asc/msi-connector/images/sha256:877a6f2a212c44021281f80cb1f4c73a09dce4e99a8cb8efcc03f7ce3c877a6f"),
		// 				},
		// 				Status: &armsecurity.SubAssessmentStatus{
		// 					Description: to.Ptr("The resource is unhealthy"),
		// 					Cause: to.Ptr(""),
		// 					Code: to.Ptr(armsecurity.SubAssessmentStatusCodeUnhealthy),
		// 					Severity: to.Ptr(armsecurity.SeverityHigh),
		// 				},
		// 				TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-23T12:20:08.7644808Z"); return t}()),
		// 			},
		// 		},
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
		// 				TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-23T12:20:08.7644808Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
