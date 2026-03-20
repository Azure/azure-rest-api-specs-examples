
/**
 * Samples for ConnectedRegistries Deactivate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ConnectedRegistryDeactivate.json
     */
    /**
     * Sample code: ConnectedRegistryDeactivate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        connectedRegistryDeactivate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getConnectedRegistries().deactivate("myResourceGroup", "myRegistry",
            "myConnectedRegistry", com.azure.core.util.Context.NONE);
    }
}
