/** Samples for ContainerAppsAuthConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/AuthConfigs_Get.json
     */
    /**
     * Sample code: Get Container App's AuthConfig.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerAppSAuthConfig(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsAuthConfigs()
            .getWithResponse("workerapps-rg-xj", "testcanadacentral", "current", com.azure.core.util.Context.NONE);
    }
}
