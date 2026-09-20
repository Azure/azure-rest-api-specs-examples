
/**
 * Samples for ContainerAppPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppPrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: Delete a Private Endpoint Connection by Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteAPrivateEndpointConnectionByContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppPrivateEndpointConnections().delete("examplerg", "testcontainerapp0",
            "test-private-endpoint-connection", com.azure.core.util.Context.NONE);
    }
}
