import com.azure.core.util.Context;

/** Samples for ContainerGroups Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsRestart.json
     */
    /**
     * Sample code: ContainerRestart.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerRestart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().restart("demo", "demo1", Context.NONE);
    }
}
