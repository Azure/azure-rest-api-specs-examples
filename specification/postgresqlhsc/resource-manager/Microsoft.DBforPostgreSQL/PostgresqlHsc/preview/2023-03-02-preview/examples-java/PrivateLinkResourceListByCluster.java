
/**
 * Samples for PrivateLinkResources ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/PrivateLinkResourceListByCluster.json
     */
    /**
     * Sample code: Gets the private link resources for cluster.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getsThePrivateLinkResourcesForCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.privateLinkResources().listByCluster("TestResourceGroup", "testcluster",
            com.azure.core.util.Context.NONE);
    }
}
