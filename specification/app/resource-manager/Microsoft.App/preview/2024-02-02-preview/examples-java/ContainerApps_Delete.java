
/**
 * Samples for ContainerApps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/ContainerApps_Delete.json
     */
    /**
     * Sample code: Delete Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().delete("rg", "testWorkerApp0", com.azure.core.util.Context.NONE);
    }
}
