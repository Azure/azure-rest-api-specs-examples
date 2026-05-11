
/**
 * Samples for QueryPacks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/QueryPacksListByResourceGroup.json
     */
    /**
     * Sample code: QueryPackListByResourceGroup.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void
        queryPackListByResourceGroup(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().listByResourceGroup("my-resource-group", com.azure.core.util.Context.NONE);
    }
}
