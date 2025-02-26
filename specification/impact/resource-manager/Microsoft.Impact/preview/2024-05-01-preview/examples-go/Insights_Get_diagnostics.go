package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/Insights_Get_diagnostics.json
func ExampleInsightsClient_Get_getInsightSampleForDiagnosticsCategory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Get(ctx, "impactid", "insight1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientGetResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactid/insights/insight1"),
	// 		Name: to.Ptr("insight1"),
	// 		Type: to.Ptr("Microsoft.impact/workloadimpacts/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			Category: to.Ptr("Diagnostics"),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactid"),
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/virtualamchine/vm"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 			},
	// 			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("We ran diagnostics on your resource and found an issue"),
	// 				Description: to.Ptr("<!--issueDescription-->\n<p>The physical host node where your VM is running had a networking stack update. This might result in a brief connectivity loss.</p>\n<br/><table><tr><th align=\"left\">Resource</th><th align=\"left\">Impact Start Time</th><th align=\"left\">Impact End Time</th><th align=\"left\">Impact Duration(Timespan)</th></tr><tr><td align=\"left\">west-eur-VM</td><td align=\"left\">2024-02-18 01:09:31 UTC</td><td align=\"left\">2024-02-18 01:09:35 UTC</td><td align=\"left\">00:00:04.2690000</td></tr></table>\n<!--/issueDescription-->\n<p>Azure performs updates to improve reliability, performance, and security of the VMs. Azure chooses the least impactful method, which might result in a brief connectivity loss. We are continuously working to improve and reduce impact of our updates, and we apologize for any inconvenience this may have caused you.</p>\n<!--recommendedActions-->\n<h2><strong>Recommended Documents</strong></h2>\n<ul>\n<li>To prepare for VM maintenance events and reduce its impact, try using <a href=\"https://docs.microsoft.com/azure/virtual-machines/windows/scheduled-events\">Scheduled Events for Windows</a> or <a href=\"https://docs.microsoft.com/azure/virtual-machines/linux/scheduled-events\">Linux</a></li>\n<li>Learn more about Azure maintenance and configuring for high availability:\n<ul>\n<li><a href=\"https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates\">Maintenance and updates for virtual machines in Azure</a></li>\n<li><a href=\"https://docs.microsoft.com/azure/virtual-machines/windows/tutorial-availability-sets\">Configure availability of virtual machines</a></li>\n</ul>\n</li>\n<li>To troubleshoot this scenario in the future, see <a href=\"https://docs.microsoft.com/azure/resource-health/resource-health-overview\">Resource Health Center</a></li>\n</ul>\n"),
	// 			},
	// 		},
	// 	},
	// }
}
