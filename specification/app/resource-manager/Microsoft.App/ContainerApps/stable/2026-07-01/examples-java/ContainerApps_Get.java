
/**
 * Samples for ContainerApps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_Get.json
     */
    /**
     * Sample code: Get Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().getByResourceGroupWithResponse("rg", "testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
