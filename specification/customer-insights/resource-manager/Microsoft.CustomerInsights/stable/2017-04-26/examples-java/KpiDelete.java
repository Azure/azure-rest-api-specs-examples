/** Samples for Kpi Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiDelete.json
     */
    /**
     * Sample code: Kpi_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void kpiDelete(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.kpis().delete("TestHubRG", "sdkTestHub", "kpiTest45453647", com.azure.core.util.Context.NONE);
    }
}
