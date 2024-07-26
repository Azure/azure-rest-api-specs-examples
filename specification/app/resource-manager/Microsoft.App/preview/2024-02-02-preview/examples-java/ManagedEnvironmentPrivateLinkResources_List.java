
/**
 * Samples for ManagedEnvironmentPrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ManagedEnvironmentPrivateLinkResources_List.json
     */
    /**
     * Sample code: List Private Link Resources by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listPrivateLinkResourcesByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentPrivateLinkResources().list("examplerg", "managedEnv",
            com.azure.core.util.Context.NONE);
    }
}
