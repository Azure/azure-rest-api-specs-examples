
/**
 * Samples for ManagedEnvironmentPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ManagedEnvironmentPrivateEndpointConnections_List.json
     */
    /**
     * Sample code: List Private Endpoint Connections by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listPrivateEndpointConnectionsByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentPrivateEndpointConnections().list("examplerg", "managedEnv",
            com.azure.core.util.Context.NONE);
    }
}
