/** Samples for ContainerGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupsDelete.json
     */
    /**
     * Sample code: ContainerGroupsDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainerGroups()
            .delete("demo", "demo1", com.azure.core.util.Context.NONE);
    }
}
