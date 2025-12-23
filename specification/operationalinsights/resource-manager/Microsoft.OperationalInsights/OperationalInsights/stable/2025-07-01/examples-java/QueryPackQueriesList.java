
/**
 * Samples for Queries List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/QueryPackQueriesList.json
     */
    /**
     * Sample code: QueryList.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queries().list("my-resource-group", "my-querypack", null, true, null, com.azure.core.util.Context.NONE);
    }
}
