
/**
 * Samples for Configurations ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/ConfigurationListByCluster.json
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
