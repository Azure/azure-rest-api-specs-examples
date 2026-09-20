
/**
 * Samples for ContainerApps Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_Start.json
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
