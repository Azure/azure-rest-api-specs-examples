/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getsPrivateEndpointConnection(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "TestGroup", "testcluster", "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
