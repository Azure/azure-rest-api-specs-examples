
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/
     * PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.privateEndpointConnections().getWithResponse("Default", "test-svr", "private-endpoint-connection-name",
            com.azure.core.util.Context.NONE);
    }
}
