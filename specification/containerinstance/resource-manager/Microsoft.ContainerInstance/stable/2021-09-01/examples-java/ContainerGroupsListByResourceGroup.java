import com.azure.core.util.Context;

/** Samples for ContainerGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsListByResourceGroup.json
     */
    /**
     * Sample code: ContainerGroupsListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainerGroups()
            .listByResourceGroup("demo", Context.NONE);
    }
}
