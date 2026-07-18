
/**
 * Samples for PrivateEndpointConnections ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsListOfPrivateEndpointConnectionsOnAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().listByServer("Default", "test-svr",
            com.azure.core.util.Context.NONE);
    }
}
