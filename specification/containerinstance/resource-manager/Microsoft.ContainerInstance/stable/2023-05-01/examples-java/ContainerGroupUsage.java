
/** Samples for Location ListUsage. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerGroupUsage.json
     */
    /**
     * Sample code: ContainerUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getLocations().listUsage("westcentralus",
            com.azure.core.util.Context.NONE);
    }
}
