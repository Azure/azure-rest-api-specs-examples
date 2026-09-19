
/**
 * Samples for LogicApps ListWorkflowsConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LogicApps_ListConnections.json
     */
    /**
     * Sample code: List the Workflows Configuration Connections.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listTheWorkflowsConfigurationConnections(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.logicApps().listWorkflowsConnectionsWithResponse("testrg123", "testapp2", "testapp2",
            com.azure.core.util.Context.NONE);
    }
}
