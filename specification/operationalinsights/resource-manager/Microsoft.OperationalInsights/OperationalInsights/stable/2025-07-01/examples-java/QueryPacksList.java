
/**
 * Samples for QueryPacks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/QueryPacksList.json
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
