
/**
 * Samples for ContainerAppsAuthConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AuthConfigs_Get.json
     */
    /**
     * Sample code: Get Container App's AuthConfig.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppSAuthConfig(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsAuthConfigs().getWithResponse("workerapps-rg-xj", "testcanadacentral", "current",
            com.azure.core.util.Context.NONE);
    }
}
