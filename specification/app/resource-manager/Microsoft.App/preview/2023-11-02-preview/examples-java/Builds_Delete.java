
/**
 * Samples for Builds Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Builds_Delete.json
     */
    /**
     * Sample code: Builds_Delete_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void buildsDelete0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.builds().delete("rg", "testBuilder", "testBuild", com.azure.core.util.Context.NONE);
    }
}
