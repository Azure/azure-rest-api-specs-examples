/** Samples for ConnectorMappings ListByConnector. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsListByConnector.json
     */
    /**
     * Sample code: ConnectorMappings_ListByConnector.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void connectorMappingsListByConnector(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .connectorMappings()
            .listByConnector("TestHubRG", "sdkTestHub", "testConnector8858", com.azure.core.util.Context.NONE);
    }
}
