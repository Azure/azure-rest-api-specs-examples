
/**
 * Samples for ContainerAppsSessionPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/SessionPools_Delete.json
     */
    /**
     * Sample code: Delete Session Pool.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteSessionPool(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSessionPools().delete("rg", "testsessionpool", com.azure.core.util.Context.NONE);
    }
}
