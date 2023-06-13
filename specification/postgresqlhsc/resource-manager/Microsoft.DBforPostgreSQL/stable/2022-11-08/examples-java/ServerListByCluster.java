/** Samples for Servers ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ServerListByCluster.json
     */
    /**
     * Sample code: List servers of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void listServersOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.servers().listByCluster("TestGroup", "testcluster1", com.azure.core.util.Context.NONE);
    }
}
