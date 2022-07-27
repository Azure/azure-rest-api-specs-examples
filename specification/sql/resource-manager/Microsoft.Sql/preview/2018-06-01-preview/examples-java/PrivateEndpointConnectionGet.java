import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .getWithResponse("Default", "test-svr", "private-endpoint-connection-name", Context.NONE);
    }
}
