
/**
 * Samples for ManagedEnvironmentPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
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
