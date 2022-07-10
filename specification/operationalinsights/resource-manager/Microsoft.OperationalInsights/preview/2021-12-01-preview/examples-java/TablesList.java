import com.azure.core.util.Context;

/** Samples for Tables ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/TablesList.json
     */
    /**
     * Sample code: TablesListByWorkspace.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesListByWorkspace(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().listByWorkspace("oiautorest6685", "oiautorest6685", Context.NONE);
    }
}
