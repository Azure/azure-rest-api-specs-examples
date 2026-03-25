
/**
 * Samples for Registries ListPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/RegistryListPrivateLinkResources.json
     */
    /**
     * Sample code: RegistryListPrivateLinkResources.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryListPrivateLinkResources(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().listPrivateLinkResources("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
