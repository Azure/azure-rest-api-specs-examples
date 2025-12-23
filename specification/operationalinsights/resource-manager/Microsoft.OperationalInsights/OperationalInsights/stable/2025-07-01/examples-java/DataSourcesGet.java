
/**
 * Samples for DataSources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/DataSourcesGet.json
     */
    /**
     * Sample code: DataSourcesGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataSourcesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataSources().getWithResponse("OIAutoRest5123", "AzTest9724", "AzTestDS774",
            com.azure.core.util.Context.NONE);
    }
}
