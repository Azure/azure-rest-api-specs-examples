
/**
 * Samples for Registries GetPrivateLinkResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryGetPrivateLinkResource.json
     */
    /**
     * Sample code: RegistryGetPrivateLinkResource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryGetPrivateLinkResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().getPrivateLinkResourceWithResponse(
            "myResourceGroup", "myRegistry", "registry", com.azure.core.util.Context.NONE);
    }
}
