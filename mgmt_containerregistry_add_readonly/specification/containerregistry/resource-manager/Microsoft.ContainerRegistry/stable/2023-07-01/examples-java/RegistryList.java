
/** Samples for Registries List. */
public final class Main {
    /*
     * x-ms-original-file:
     * mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/
     * stable/2023-07-01/examples/RegistryList.json
     */
    /**
     * Sample code: RegistryList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().list(com.azure.core.util.Context.NONE);
    }
}
