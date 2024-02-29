
/**
 * Samples for Registries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/
     * RegistryDelete.json
     */
    /**
     * Sample code: RegistryDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().delete("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
