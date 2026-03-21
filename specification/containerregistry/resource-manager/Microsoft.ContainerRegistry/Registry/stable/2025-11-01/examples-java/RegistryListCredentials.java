
/**
 * Samples for Registries ListCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryListCredentials.json
     */
    /**
     * Sample code: RegistryListCredentials.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryListCredentials(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().listCredentialsWithResponse("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
