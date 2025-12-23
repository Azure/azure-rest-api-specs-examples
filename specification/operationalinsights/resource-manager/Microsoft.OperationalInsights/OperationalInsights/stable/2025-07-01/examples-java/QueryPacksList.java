
/**
 * Samples for QueryPacks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/QueryPacksList.json
     */
    /**
     * Sample code: QueryPacksList.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPacksList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().list(com.azure.core.util.Context.NONE);
    }
}
