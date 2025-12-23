
import com.azure.resourcemanager.loganalytics.models.WorkspacePurgeBody;
import com.azure.resourcemanager.loganalytics.models.WorkspacePurgeBodyFilters;
import java.util.Arrays;

/**
 * Samples for WorkspacePurge Purge.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesPurge.json
     */
    /**
     * Sample code: WorkspacePurge.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacePurge(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspacePurges().purgeWithResponse("OIAutoRest5123", "aztest5048",
            new WorkspacePurgeBody().withTable("Heartbeat").withFilters(Arrays.asList(new WorkspacePurgeBodyFilters()
                .withColumn("TimeGenerated").withOperator(">").withValue("2017-09-01T00:00:00"))),
            com.azure.core.util.Context.NONE);
    }
}
