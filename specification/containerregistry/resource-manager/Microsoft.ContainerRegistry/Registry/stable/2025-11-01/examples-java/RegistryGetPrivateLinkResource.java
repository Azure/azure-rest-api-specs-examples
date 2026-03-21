
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
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryGetPrivateLinkResource(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().getPrivateLinkResourceWithResponse("myResourceGroup", "myRegistry",
            "registry", com.azure.core.util.Context.NONE);
    }
}
