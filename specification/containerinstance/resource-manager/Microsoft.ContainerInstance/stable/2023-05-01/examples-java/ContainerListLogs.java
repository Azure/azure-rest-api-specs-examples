
/** Samples for Containers ListLogs. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerListLogs.json
     */
    /**
     * Sample code: ContainerListLogs.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerListLogs(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainers().listLogsWithResponse("demo", "demo1",
            "container1", 10, null, com.azure.core.util.Context.NONE);
    }
}
