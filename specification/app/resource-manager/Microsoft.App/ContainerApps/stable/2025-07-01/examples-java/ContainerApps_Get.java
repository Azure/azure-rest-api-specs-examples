
/**
 * Samples for ContainerAppsDiagnostics GetRoot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ContainerApps_Get.json
     */
    /**
     * Sample code: Get Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsDiagnostics().getRootWithResponse("rg", "testcontainerapp0",
            com.azure.core.util.Context.NONE);
    }
}
