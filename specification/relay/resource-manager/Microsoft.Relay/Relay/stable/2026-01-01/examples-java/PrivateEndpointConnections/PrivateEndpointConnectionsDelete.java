
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateEndpointConnections/PrivateEndpointConnectionsDelete.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void nameSpacePrivateEndPointConnectionDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.privateEndpointConnections().delete("myResourceGroup", "example-RelayNamespace-5849",
            "{privateEndpointConnection name}", com.azure.core.util.Context.NONE);
    }
}
