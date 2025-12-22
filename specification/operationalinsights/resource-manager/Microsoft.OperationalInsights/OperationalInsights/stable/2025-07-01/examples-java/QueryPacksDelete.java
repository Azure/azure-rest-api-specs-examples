
/**
 * Samples for QueryPacks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/QueryPacksDelete.json
     */
    /**
     * Sample code: QueryPacksDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPacksDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().deleteByResourceGroupWithResponse("my-resource-group", "my-querypack",
            com.azure.core.util.Context.NONE);
    }
}
