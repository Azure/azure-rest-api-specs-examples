/** Samples for QueryPacks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksList.json
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
