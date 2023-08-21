/** Samples for ContainerApps Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerApps_Stop.json
     */
    /**
     * Sample code: Stop Container App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void stopContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().stop("rg", "testWorkerApp0", com.azure.core.util.Context.NONE);
    }
}
