
/**
 * Samples for Registries ListPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryListPrivateLinkResources.json
     */
    /**
     * Sample code: RegistryListPrivateLinkResources.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries()
            .listPrivateLinkResources("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
