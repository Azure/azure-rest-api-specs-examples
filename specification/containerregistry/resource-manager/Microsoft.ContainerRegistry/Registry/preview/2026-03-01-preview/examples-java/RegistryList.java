
/**
 * Samples for Registries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/RegistryList.json
     */
    /**
     * Sample code: RegistryList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void registryList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().list(com.azure.core.util.Context.NONE);
    }
}
