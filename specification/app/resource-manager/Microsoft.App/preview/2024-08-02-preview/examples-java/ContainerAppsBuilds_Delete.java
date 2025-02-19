
/**
 * Samples for ContainerAppsBuilds Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsBuilds_Delete.
     * json
     */
    /**
     * Sample code: ContainerAppsBuilds_Delete_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        containerAppsBuildsDelete0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsBuilds().delete("rg", "testCapp", "testBuild", com.azure.core.util.Context.NONE);
    }
}
