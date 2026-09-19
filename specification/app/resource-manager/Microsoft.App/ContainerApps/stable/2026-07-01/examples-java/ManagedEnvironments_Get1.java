
/**
 * Samples for ManagedEnvironmentsDiagnostics GetRoot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironments_Get1.json
     */
    /**
     * Sample code: Get environments by name.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getEnvironmentsByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentsDiagnostics().getRootWithResponse("examplerg", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
