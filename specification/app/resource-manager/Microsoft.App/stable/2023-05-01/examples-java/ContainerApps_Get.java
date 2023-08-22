/** Samples for ContainerAppsDiagnostics GetRoot. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerApps_Get.json
     */
    /**
     * Sample code: Get Container App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsDiagnostics()
            .getRootWithResponse("rg", "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
