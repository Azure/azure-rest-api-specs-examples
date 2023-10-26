/** Samples for DataSources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataSourcesDelete.json
     */
    /**
     * Sample code: DataSourcesDelete.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataSourcesDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .dataSources()
            .deleteWithResponse("OIAutoRest5123", "AzTest9724", "AzTestDS774", com.azure.core.util.Context.NONE);
    }
}
