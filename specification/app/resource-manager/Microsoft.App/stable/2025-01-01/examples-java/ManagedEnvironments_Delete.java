
/**
 * Samples for ManagedEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ManagedEnvironments_Delete.json
     */
    /**
     * Sample code: Delete environment by name.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        deleteEnvironmentByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().delete("examplerg", "examplekenv", com.azure.core.util.Context.NONE);
    }
}
