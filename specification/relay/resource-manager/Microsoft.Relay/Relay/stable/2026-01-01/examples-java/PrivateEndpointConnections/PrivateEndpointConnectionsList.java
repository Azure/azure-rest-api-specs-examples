
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateEndpointConnections/PrivateEndpointConnectionsList.json
     */
    /**
     * Sample code: PrivateEndpointConnectionsList.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void privateEndpointConnectionsList(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.privateEndpointConnections().list("myResourceGroup", "example-RelayNamespace-5849",
            com.azure.core.util.Context.NONE);
    }
}
