
/**
 * Samples for Servers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * ServerGet.json
     */
    /**
     * Sample code: Get the server of cluster.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void
        getTheServerOfCluster(com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.servers().getWithResponse("TestGroup", "testcluster1", "testcluster1-c",
            com.azure.core.util.Context.NONE);
    }
}
