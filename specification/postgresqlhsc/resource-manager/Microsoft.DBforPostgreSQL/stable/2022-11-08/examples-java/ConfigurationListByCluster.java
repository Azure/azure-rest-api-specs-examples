/** Samples for Configurations ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationListByCluster.json
     */
    /**
     * Sample code: List configurations of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void listConfigurationsOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.configurations().listByCluster("TestResourceGroup", "testcluster", com.azure.core.util.Context.NONE);
    }
}
