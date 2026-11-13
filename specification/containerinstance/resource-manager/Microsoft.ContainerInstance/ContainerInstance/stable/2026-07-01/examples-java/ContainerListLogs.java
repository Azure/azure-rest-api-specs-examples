
/**
 * Samples for Containers ListLogs.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerListLogs.json
     */
    /**
     * Sample code: ContainerListLogs.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerListLogs(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainers().listLogsWithResponse("demo", "demo1", "container1", 10, null,
            com.azure.core.util.Context.NONE);
    }
}
