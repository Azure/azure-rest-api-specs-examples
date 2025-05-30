
/**
 * Samples for ContainerAppsPatches Apply.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsPatches_Apply.
     * json
     */
    /**
     * Sample code: ContainerAppsPatches_Apply_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        containerAppsPatchesApply0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsPatches().apply("rg", "test-app", "testPatch-25fe4b", com.azure.core.util.Context.NONE);
    }
}
