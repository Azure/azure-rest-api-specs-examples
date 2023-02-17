/** Samples for Registries GetPrivateLinkResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/RegistryGetPrivateLinkResource.json
     */
    /**
     * Sample code: RegistryGetPrivateLinkResource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryGetPrivateLinkResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .getPrivateLinkResourceWithResponse(
                "myResourceGroup", "myRegistry", "registry", com.azure.core.util.Context.NONE);
    }
}
