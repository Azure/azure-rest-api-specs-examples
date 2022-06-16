import com.azure.core.util.Context;

/** Samples for Registries ListUsages. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/RegistryListUsages.json
     */
    /**
     * Sample code: RegistryListUsages.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .listUsagesWithResponse("myResourceGroup", "myRegistry", Context.NONE);
    }
}
