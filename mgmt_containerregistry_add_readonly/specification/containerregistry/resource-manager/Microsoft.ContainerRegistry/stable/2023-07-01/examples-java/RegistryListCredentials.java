/** Samples for Registries ListCredentials. */
public final class Main {
    /*
     * x-ms-original-file: mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/RegistryListCredentials.json
     */
    /**
     * Sample code: RegistryListCredentials.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListCredentials(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .listCredentialsWithResponse("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
