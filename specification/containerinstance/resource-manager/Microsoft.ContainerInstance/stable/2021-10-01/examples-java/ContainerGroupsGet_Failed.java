import com.azure.core.util.Context;

/** Samples for ContainerGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsGet_Failed.json
     */
    /**
     * Sample code: ContainerGroupsGet_Failed.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsGetFailed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainerGroups()
            .getByResourceGroupWithResponse("demo", "demo1", Context.NONE);
    }
}
