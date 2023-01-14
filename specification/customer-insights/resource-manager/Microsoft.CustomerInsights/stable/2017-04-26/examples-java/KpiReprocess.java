/** Samples for Kpi Reprocess. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiReprocess.json
     */
    /**
     * Sample code: Kpi_Reprocess.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void kpiReprocess(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .kpis()
            .reprocessWithResponse("TestHubRG", "sdkTestHub", "kpiTest45453647", com.azure.core.util.Context.NONE);
    }
}
