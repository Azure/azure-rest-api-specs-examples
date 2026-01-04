
/**
 * Samples for Replications List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * ReplicationList.json
     */
    /**
     * Sample code: ReplicationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getReplications().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
