
/**
 * Samples for PrivateEndpointConnectionsInterface ListByTrafficController.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Get Private Endpoint Connections.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getPrivateEndpointConnections(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.privateEndpointConnectionsInterfaces().listByTrafficController("rg1", "tc1",
            com.azure.core.util.Context.NONE);
    }
}
