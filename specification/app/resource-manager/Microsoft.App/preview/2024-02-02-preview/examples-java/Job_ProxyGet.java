
/**
 * Samples for Jobs ProxyGet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Job_ProxyGet.json
     */
    /**
     * Sample code: Get Container App Job by name.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppJobByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().proxyGetWithResponse("rg", "testcontainerAppsJob0", com.azure.core.util.Context.NONE);
    }
}
