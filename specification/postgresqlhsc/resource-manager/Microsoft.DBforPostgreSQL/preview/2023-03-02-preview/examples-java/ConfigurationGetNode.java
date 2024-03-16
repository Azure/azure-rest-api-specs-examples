
/**
 * Samples for Configurations GetNode.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * ConfigurationGetNode.json
     */
    /**
     * Sample code: Get configuration details for node.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getConfigurationDetailsForNode(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.configurations().getNodeWithResponse("TestResourceGroup", "testcluster", "array_nulls",
            com.azure.core.util.Context.NONE);
    }
}
