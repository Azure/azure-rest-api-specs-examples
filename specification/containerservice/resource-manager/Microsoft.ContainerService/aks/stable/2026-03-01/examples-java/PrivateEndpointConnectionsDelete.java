
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateEndpointConnectionsDelete.json
     */
    /**
     * Sample code: Delete Private Endpoint Connection.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deletePrivateEndpointConnection(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().delete("rg1", "clustername1",
            "privateendpointconnection1", com.azure.core.util.Context.NONE);
    }
}
