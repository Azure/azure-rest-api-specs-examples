
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
     * 07-01/examples/WorkspacesPurgeResourceId.json
     */
    /**
     * Sample code: WorkspacePurgeResourceId.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacePurgeResourceId(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspacePurges().purgeWithResponse("OIAutoRest5123", "aztest5048",
            new WorkspacePurgeBody().withTable("Heartbeat").withFilters(
                Arrays.asList(new WorkspacePurgeBodyFilters().withColumn("_ResourceId").withOperator("==").withValue(
                    "/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/SomeResourceGroup/providers/microsoft.insights/components/AppInsightResource"))),
            com.azure.core.util.Context.NONE);
    }
}
