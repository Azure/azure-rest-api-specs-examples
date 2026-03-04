
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/PrivateEndpointConnectionsDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void deletesAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.privateEndpointConnections().delete("TestGroup", "testcluster", "private-endpoint-connection-name",
            com.azure.core.util.Context.NONE);
    }
}
