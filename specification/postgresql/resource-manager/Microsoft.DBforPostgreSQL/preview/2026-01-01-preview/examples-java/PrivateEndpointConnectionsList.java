
/**
 * Samples for PrivateEndpointConnections ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/PrivateEndpointConnectionsList.json
     */
    /**
     * Sample code: List all private endpoint connections on a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllPrivateEndpointConnectionsOnAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.privateEndpointConnections().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
