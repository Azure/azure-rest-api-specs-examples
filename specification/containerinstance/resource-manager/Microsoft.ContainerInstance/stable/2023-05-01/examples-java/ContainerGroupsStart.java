
/**
 * Samples for ContainerGroups Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerGroupsStart.json
     */
    /**
     * Sample code: ContainerStart.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().start("demo", "demo1",
            com.azure.core.util.Context.NONE);
    }
}
