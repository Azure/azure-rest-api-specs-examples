
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * ClusterDelete.json
     */
    /**
     * Sample code: Delete the cluster.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void
        deleteTheCluster(com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().delete("TestGroup", "testcluster", com.azure.core.util.Context.NONE);
    }
}
