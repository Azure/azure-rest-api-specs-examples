/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "Default", "test-svr", "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
