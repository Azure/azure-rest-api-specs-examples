
/**
 * Samples for PrivateEndpointConnections ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/PrivateEndpointConnectionsListByCluster.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a cluster.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getsListOfPrivateEndpointConnectionsOnACluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.privateEndpointConnections().listByCluster("TestResourceGroup", "testcluster",
            com.azure.core.util.Context.NONE);
    }
}
