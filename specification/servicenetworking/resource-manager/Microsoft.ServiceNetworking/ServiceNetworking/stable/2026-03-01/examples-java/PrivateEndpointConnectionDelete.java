
/**
 * Samples for PrivateEndpointConnectionsInterface Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Delete Private Endpoint Connection.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        deletePrivateEndpointConnection(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.privateEndpointConnectionsInterfaces().delete("rg1", "tc1", "pec1", com.azure.core.util.Context.NONE);
    }
}
