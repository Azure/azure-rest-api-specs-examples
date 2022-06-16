import com.azure.core.util.Context;

/** Samples for Registries ListPrivateLinkResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/RegistryListPrivateLinkResources.json
     */
    /**
     * Sample code: RegistryListPrivateLinkResources.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .listPrivateLinkResources("myResourceGroup", "myRegistry", Context.NONE);
    }
}
