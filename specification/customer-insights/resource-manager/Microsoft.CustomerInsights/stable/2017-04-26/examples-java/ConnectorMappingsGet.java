/** Samples for ConnectorMappings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsGet.json
     */
    /**
     * Sample code: ConnectorMappings_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void connectorMappingsGet(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .connectorMappings()
            .getWithResponse(
                "TestHubRG", "sdkTestHub", "testConnector8858", "testMapping12491", com.azure.core.util.Context.NONE);
    }
}
