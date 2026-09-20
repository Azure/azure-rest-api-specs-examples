
/**
 * Samples for ContainerAppsSourceControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SourceControls_Get.json
     */
    /**
     * Sample code: Get Container App's SourceControl.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppSSourceControl(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSourceControls().getWithResponse("workerapps-rg-xj", "testcanadacentral", "current",
            com.azure.core.util.Context.NONE);
    }
}
