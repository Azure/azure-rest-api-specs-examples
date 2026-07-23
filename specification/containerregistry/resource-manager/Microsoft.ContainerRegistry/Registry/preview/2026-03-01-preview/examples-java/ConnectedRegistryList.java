
/**
 * Samples for ConnectedRegistries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ConnectedRegistryList.json
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
