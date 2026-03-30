
/**
 * Samples for WebApps ListWorkflowsConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWorkflowsConfigurationConnections.json
     */
    /**
     * Sample code: List the Instance Workflows Configuration Connections.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listTheInstanceWorkflowsConfigurationConnections(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listWorkflowsConnectionsWithResponse("testrg123", "testsite2",
            com.azure.core.util.Context.NONE);
    }
}
