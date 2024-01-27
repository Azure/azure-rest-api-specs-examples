
/** Samples for Containers Attach. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerAttach.json
     */
    /**
     * Sample code: ContainerAttach.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerAttach(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainers().attachWithResponse("demo", "demo1",
            "container1", com.azure.core.util.Context.NONE);
    }
}
