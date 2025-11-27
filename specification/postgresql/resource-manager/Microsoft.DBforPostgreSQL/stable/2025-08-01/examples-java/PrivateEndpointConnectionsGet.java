
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Get a private endpoint connection.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getAPrivateEndpointConnection(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.privateEndpointConnections().getWithResponse("exampleresourcegroup", "exampleserver",
            "private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e", com.azure.core.util.Context.NONE);
    }
}
