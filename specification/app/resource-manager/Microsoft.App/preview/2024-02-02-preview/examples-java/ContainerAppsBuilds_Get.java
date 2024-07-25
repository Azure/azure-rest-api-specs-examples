
/**
 * Samples for ContainerAppsBuilds Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/ContainerAppsBuilds_Get.json
     */
    /**
     * Sample code: ContainerAppsBuilds_Get_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        containerAppsBuildsGet0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsBuilds().getWithResponse("rg", "testCapp", "testBuild", com.azure.core.util.Context.NONE);
    }
}
