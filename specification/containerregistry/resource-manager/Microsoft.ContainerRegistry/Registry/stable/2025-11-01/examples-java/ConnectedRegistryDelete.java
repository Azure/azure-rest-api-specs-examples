
/**
 * Samples for ConnectedRegistries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * ConnectedRegistryDelete.json
     */
    /**
     * Sample code: ConnectedRegistryDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectedRegistryDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getConnectedRegistries().delete("myResourceGroup",
            "myRegistry", "myConnectedRegistry", com.azure.core.util.Context.NONE);
    }
}
