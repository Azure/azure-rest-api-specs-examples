
/**
 * Samples for ContainerAppsDiagnostics GetRoot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_Get1.json
     */
    /**
     * Sample code: Get Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsDiagnostics().getRootWithResponse("rg", "testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
