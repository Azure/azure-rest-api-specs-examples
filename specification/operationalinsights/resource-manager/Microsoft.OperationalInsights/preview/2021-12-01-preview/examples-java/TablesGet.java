import com.azure.core.util.Context;

/** Samples for Tables Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/TablesGet.json
     */
    /**
     * Sample code: TablesGet.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void tablesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.tables().getWithResponse("oiautorest6685", "oiautorest6685", "table1_SRCH", Context.NONE);
    }
}
