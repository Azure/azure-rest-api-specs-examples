
/**
 * Samples for Registries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryDelete.json
     */
    /**
     * Sample code: RegistryDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void registryDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().delete("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
