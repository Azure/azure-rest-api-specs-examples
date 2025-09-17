
/**
 * Samples for ManagedEnvironmentsDiagnostics GetRoot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * ManagedEnvironments_Get.json
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
