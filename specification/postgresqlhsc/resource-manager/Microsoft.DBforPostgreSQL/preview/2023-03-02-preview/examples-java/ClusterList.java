
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * ClusterList.json
     */
    /**
     * Sample code: List all the clusters.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void
        listAllTheClusters(com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
