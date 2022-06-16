import com.azure.core.util.Context;

/** Samples for Registries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/RegistryDelete.json
     */
    /**
     * Sample code: RegistryDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .delete("myResourceGroup", "myRegistry", Context.NONE);
    }
}
