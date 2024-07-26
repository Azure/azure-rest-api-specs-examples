
/**
 * Samples for ManagedEnvironmentPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ManagedEnvironmentPrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: Delete a Private Endpoint Connection by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteAPrivateEndpointConnectionByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentPrivateEndpointConnections().delete("examplerg", "managedEnv", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
