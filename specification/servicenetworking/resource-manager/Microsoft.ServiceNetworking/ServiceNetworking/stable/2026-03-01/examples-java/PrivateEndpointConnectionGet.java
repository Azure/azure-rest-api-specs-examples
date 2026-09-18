
/**
 * Samples for PrivateEndpointConnectionsInterface Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Get Private Endpoint Connection.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getPrivateEndpointConnection(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.privateEndpointConnectionsInterfaces().getWithResponse("rg1", "tc1", "pec1",
            com.azure.core.util.Context.NONE);
    }
}
