
/**
 * Samples for ConnectedRegistries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/ConnectedRegistryDelete.json
     */
    /**
     * Sample code: ConnectedRegistryDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        connectedRegistryDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getConnectedRegistries().delete("myResourceGroup", "myRegistry", "myConnectedRegistry",
            com.azure.core.util.Context.NONE);
    }
}
