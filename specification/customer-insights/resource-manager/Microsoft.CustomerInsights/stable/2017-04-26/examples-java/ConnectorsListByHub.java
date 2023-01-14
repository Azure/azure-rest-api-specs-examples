/** Samples for Connectors ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorsListByHub.json
     */
    /**
     * Sample code: Connectors_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void connectorsListByHub(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.connectors().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
