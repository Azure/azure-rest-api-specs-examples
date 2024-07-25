
/**
 * Samples for Builds Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Builds_Get.json
     */
    /**
     * Sample code: Builds_Get_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void buildsGet0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.builds().getWithResponse("rg", "testBuilder", "testBuild", com.azure.core.util.Context.NONE);
    }
}
