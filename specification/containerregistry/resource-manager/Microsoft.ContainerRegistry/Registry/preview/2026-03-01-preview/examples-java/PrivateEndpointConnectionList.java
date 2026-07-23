
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: PrivateEndpointConnectionList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        privateEndpointConnectionList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
