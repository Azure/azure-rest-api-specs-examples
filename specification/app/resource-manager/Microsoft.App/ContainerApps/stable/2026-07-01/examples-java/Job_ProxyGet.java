
/**
 * Samples for Jobs ProxyGet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_ProxyGet.json
     */
    /**
     * Sample code: Get Container App Job by name.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppJobByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().proxyGetWithResponse("rg", "testcontainerAppsJob0", "rootApi", com.azure.core.util.Context.NONE);
    }
}
