
/**
 * Samples for Registries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * RegistryListByResourceGroup.json
     */
    /**
     * Sample code: RegistryListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
