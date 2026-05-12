
/**
 * Samples for DataSources ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DataSourcesListByWorkspace.json
     */
    /**
     * Sample code: DataSourcesListByWorkspace.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataSourcesListByWorkspace(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataSources().listByWorkspace("OIAutoRest5123", "AzTest9724", "kind='WindowsEvent'", null,
            com.azure.core.util.Context.NONE);
    }
}
