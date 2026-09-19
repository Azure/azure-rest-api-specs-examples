
/**
 * Samples for ManagedEnvironmentPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentPrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: Get a Private Endpoint Connection by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getAPrivateEndpointConnectionByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentPrivateEndpointConnections().getWithResponse("examplerg", "managedEnv", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
