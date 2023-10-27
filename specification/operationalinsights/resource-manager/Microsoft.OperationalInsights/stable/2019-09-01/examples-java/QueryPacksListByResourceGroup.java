/** Samples for QueryPacks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksListByResourceGroup.json
     */
    /**
     * Sample code: QueryPackListByResourceGroup.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPackListByResourceGroup(
        com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().listByResourceGroup("my-resource-group", com.azure.core.util.Context.NONE);
    }
}
