
/**
 * Samples for ContainerAppsSessionPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SessionPools_Delete.json
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
