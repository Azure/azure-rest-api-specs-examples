/** Samples for Kpi Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiGet.json
     */
    /**
     * Sample code: Kpi_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void kpiGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.kpis().getWithResponse("TestHubRG", "sdkTestHub", "kpiTest45453647", com.azure.core.util.Context.NONE);
    }
}
