
/**
 * Samples for Registries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/RegistryListByResourceGroup.json
     */
    /**
     * Sample code: RegistryListByResourceGroup.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryListByResourceGroup(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
