import com.azure.core.util.Context;

/** Samples for ContainerGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsList.json
     */
    /**
     * Sample code: ContainerGroupsList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().list(Context.NONE);
    }
}
