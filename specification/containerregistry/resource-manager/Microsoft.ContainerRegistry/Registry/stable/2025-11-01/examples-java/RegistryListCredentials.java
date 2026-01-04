
/**
 * Samples for Registries ListCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * RegistryListCredentials.json
     */
    /**
     * Sample code: RegistryListCredentials.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListCredentials(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries()
            .listCredentialsWithResponse("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
