
/**
 * Samples for ManagedEnvironmentPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentPrivateEndpointConnections_List.json
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
