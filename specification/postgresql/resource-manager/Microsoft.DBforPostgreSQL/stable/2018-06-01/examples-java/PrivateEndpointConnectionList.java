
/**
 * Samples for PrivateEndpointConnections ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/
     * PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getsListOfPrivateEndpointConnectionsOnAServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.privateEndpointConnections().listByServer("Default", "test-svr", com.azure.core.util.Context.NONE);
    }
}
