
/**
 * Samples for Registries GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/RegistryGet.json
     */
    /**
     * Sample code: RegistryGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void registryGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().getByResourceGroupWithResponse("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
