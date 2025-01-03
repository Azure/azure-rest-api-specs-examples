
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/
     * PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     * 
     * @param manager Entry point to MariaDBManager.
     */
    public static void
        deletesAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.privateEndpointConnections().delete("Default", "test-svr", "private-endpoint-connection-name",
            com.azure.core.util.Context.NONE);
    }
}
