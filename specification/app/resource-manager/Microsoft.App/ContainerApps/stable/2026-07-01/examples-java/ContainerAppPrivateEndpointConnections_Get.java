
/**
 * Samples for ContainerAppPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppPrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: Get a Private Endpoint Connection by Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getAPrivateEndpointConnectionByContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppPrivateEndpointConnections().getWithResponse("examplerg", "testcontainerapp0",
            "test-private-endpoint-connection", com.azure.core.util.Context.NONE);
    }
}
