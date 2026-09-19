
/**
 * Samples for ContainerAppsFunctions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppsFunctions_Get.json
     */
    /**
     * Sample code: Get Container App's function.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppSFunction(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsFunctions().getWithResponse("myResourceGroup", "myContainerApp", "HttpExample",
            com.azure.core.util.Context.NONE);
    }
}
