
/**
 * Samples for ContainerApps Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ContainerApps_Start.json
     */
    /**
     * Sample code: Start Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void startContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().start("rg", "testworkerapp0", com.azure.core.util.Context.NONE);
    }
}
