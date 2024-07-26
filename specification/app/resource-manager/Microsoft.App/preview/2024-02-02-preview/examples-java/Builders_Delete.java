
/**
 * Samples for Builders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Builders_Delete.json
     */
    /**
     * Sample code: Builders_Delete_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void buildersDelete0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.builders().delete("rg", "testBuilder", com.azure.core.util.Context.NONE);
    }
}
