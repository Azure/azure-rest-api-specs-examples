
/**
 * Samples for ContainerApps Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ContainerApps_Start.json
     */
    /**
     * Sample code: Start Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void startContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().start("rg", "testWorkerApp0", com.azure.core.util.Context.NONE);
    }
}
