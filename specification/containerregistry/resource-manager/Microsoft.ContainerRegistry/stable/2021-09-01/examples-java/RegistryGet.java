import com.azure.core.util.Context;

/** Samples for Registries GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/RegistryGet.json
     */
    /**
     * Sample code: RegistryGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .getByResourceGroupWithResponse("myResourceGroup", "myRegistry", Context.NONE);
    }
}
