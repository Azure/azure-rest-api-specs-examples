
/**
 * Samples for ManagedEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironments_Delete.json
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
