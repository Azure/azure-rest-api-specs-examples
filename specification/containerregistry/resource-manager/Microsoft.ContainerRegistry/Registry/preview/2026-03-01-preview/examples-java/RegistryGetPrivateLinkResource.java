
/**
 * Samples for Registries GetPrivateLinkResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/RegistryGetPrivateLinkResource.json
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
