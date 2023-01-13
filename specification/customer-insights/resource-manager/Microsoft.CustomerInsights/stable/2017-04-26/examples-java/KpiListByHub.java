/** Samples for Kpi ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiListByHub.json
     */
    /**
     * Sample code: Kpi_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void kpiListByHub(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.kpis().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
