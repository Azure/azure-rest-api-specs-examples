
/** Samples for ContainerGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerGroupsListByResourceGroup.json
     */
    /**
     * Sample code: ContainerGroupsListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().listByResourceGroup("demo",
            com.azure.core.util.Context.NONE);
    }
}
