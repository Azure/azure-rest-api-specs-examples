
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getPrivateEndpointConnections().getWithResponse("Default",
            "test-svr", "private-endpoint-connection-name", Context.NONE);
    }
}
