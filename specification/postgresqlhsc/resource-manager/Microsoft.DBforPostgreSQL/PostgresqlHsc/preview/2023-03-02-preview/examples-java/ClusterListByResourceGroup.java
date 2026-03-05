
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/ClusterListByResourceGroup.json
     */
    /**
     * Sample code: List the clusters by resource group.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void listTheClustersByResourceGroup(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().listByResourceGroup("TestGroup", com.azure.core.util.Context.NONE);
    }
}
