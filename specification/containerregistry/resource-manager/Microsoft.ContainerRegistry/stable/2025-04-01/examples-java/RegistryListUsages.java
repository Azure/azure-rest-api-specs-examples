
/**
 * Samples for Registries ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * RegistryListUsages.json
     */
    /**
     * Sample code: RegistryListUsages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryListUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().listUsagesWithResponse("myResourceGroup",
            "myRegistry", com.azure.core.util.Context.NONE);
    }
}
