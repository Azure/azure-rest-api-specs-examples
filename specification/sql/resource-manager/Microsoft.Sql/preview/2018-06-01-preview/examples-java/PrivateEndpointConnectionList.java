import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsListOfPrivateEndpointConnectionsOnAServer(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .listByServer("Default", "test-svr", Context.NONE);
    }
}
