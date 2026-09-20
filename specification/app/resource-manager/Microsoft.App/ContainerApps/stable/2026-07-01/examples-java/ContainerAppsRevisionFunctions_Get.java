
/**
 * Samples for ContainerAppsRevisionFunctions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppsRevisionFunctions_Get.json
     */
    /**
     * Sample code: Get Container App Revision's function.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppRevisionSFunction(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisionFunctions().getWithResponse("myResourceGroup", "myContainerApp",
            "myContainerApp-abc123", "HttpExample", com.azure.core.util.Context.NONE);
    }
}
