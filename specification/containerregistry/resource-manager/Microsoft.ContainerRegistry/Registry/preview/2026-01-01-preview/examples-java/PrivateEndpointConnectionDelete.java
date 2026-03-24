
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: PrivateEndpointConnectionDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        privateEndpointConnectionDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().delete("myResourceGroup", "myRegistry", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
