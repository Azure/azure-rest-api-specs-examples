
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Get Private Endpoint Connection.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getPrivateEndpointConnection(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("rg1", "clustername1",
            "privateendpointconnection1", com.azure.core.util.Context.NONE);
    }
}
