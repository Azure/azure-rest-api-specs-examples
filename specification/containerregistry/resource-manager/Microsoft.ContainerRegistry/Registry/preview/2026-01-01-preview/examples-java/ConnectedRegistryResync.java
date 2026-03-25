
/**
 * Samples for ConnectedRegistries Resync.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ConnectedRegistryResync.json
     */
    /**
     * Sample code: ConnectedRegistryResync.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        connectedRegistryResync(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getConnectedRegistries().resyncWithResponse("myResourceGroup", "myRegistry",
            "myConnectedRegistry", com.azure.core.util.Context.NONE);
    }
}
