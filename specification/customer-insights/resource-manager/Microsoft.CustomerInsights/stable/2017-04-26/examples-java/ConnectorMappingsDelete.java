/** Samples for ConnectorMappings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsDelete.json
     */
    /**
     * Sample code: ConnectorMappings_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void connectorMappingsDelete(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .connectorMappings()
            .deleteWithResponse(
                "TestHubRG", "sdkTestHub", "testConnector8858", "testMapping12491", com.azure.core.util.Context.NONE);
    }
}
