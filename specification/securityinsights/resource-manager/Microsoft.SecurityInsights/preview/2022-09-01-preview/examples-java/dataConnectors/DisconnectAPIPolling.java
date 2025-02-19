
/**
 * Samples for DataConnectors Disconnect.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * dataConnectors/DisconnectAPIPolling.json
     */
    /**
     * Sample code: Disconnect an APIPolling data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void disconnectAnAPIPollingDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().disconnectWithResponse("myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8",
            com.azure.core.util.Context.NONE);
    }
}
