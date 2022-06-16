import com.azure.core.util.Context;

/** Samples for Registries ListCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/RegistryListCredentials.json
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
            .listCredentialsWithResponse("myResourceGroup", "myRegistry", Context.NONE);
    }
}
