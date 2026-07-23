
/**
 * Samples for Registries ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/RegistryListUsages.json
     */
    /**
     * Sample code: RegistryListUsages.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryListUsages(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().listUsagesWithResponse("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
