
/**
 * Samples for ContainerAppsPatches Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/ContainerAppsPatches_Get.
     * json
     */
    /**
     * Sample code: ContainerAppsPatches_Get_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        containerAppsPatchesGet0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsPatches().getWithResponse("rg", "test-app", "testPatch-25fe4b",
            com.azure.core.util.Context.NONE);
    }
}
