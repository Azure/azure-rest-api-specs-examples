import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("Default", "test-svr", "private-endpoint-connection-name", Context.NONE);
    }
}
