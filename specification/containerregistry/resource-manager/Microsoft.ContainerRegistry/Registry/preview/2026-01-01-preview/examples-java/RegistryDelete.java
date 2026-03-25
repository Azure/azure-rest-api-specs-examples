
/**
 * Samples for Registries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/RegistryDelete.json
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
