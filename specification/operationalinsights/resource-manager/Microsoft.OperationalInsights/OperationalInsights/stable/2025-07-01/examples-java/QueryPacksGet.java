
/**
 * Samples for QueryPacks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/QueryPacksGet.json
     */
    /**
     * Sample code: QueryPackGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPackGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().getByResourceGroupWithResponse("my-resource-group", "my-querypack",
            com.azure.core.util.Context.NONE);
    }
}
