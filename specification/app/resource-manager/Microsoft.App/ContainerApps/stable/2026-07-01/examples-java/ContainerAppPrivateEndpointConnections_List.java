
/**
 * Samples for ContainerAppPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppPrivateEndpointConnections_List.json
     */
    /**
     * Sample code: List Private Endpoint Connections by Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listPrivateEndpointConnectionsByContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppPrivateEndpointConnections().list("examplerg", "testcontainerapp0",
            com.azure.core.util.Context.NONE);
    }
}
