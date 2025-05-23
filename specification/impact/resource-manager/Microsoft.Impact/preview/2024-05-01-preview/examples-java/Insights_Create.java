
import com.azure.resourcemanager.impactreporting.models.Content;
import com.azure.resourcemanager.impactreporting.models.ImpactDetails;
import com.azure.resourcemanager.impactreporting.models.InsightProperties;
import java.time.OffsetDateTime;

/**
 * Samples for Insights Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Insights_Create.json
     */
    /**
     * Sample code: Creating an insight.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void creatingAnInsight(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.insights().define("insightId12").withExistingWorkloadImpact("impactid22")
            .withProperties(new InsightProperties().withCategory("repair").withStatus("resolved")
                .withContent(new Content().withTitle("Impact Has been correlated to an outage").withDescription(
                    "At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"))
                .withEventTime(OffsetDateTime.parse("2023-06-15T04:00:00.009223Z"))
                .withInsightUniqueId("00000000-0000-0000-0000-000000000000")
                .withImpact(new ImpactDetails().withImpactedResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername")
                    .withStartTime(OffsetDateTime.parse("2023-06-15T01:00:00.009223Z")).withImpactId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22")))
            .create();
    }
}
