
/**
 * Samples for ContainerAppsSessionPools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/SessionPools_Get.json
     */
    /**
     * Sample code: Get Session Pool.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getSessionPool(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSessionPools().getByResourceGroupWithResponse("rg", "testsessionpool",
            com.azure.core.util.Context.NONE);
    }
}
