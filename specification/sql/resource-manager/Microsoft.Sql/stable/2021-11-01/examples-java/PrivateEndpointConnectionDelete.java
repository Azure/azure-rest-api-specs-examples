
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deletesAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getPrivateEndpointConnections().delete("Default", "test-svr",
            "private-endpoint-connection-name", Context.NONE);
    }
}
