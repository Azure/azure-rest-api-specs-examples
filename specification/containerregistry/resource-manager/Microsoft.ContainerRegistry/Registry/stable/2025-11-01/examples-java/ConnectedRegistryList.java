
/**
 * Samples for ConnectedRegistries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ConnectedRegistryList.json
     */
    /**
     * Sample code: ConnectedRegistryList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        connectedRegistryList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getConnectedRegistries().list("myResourceGroup", "myRegistry", null,
            com.azure.core.util.Context.NONE);
    }
}
