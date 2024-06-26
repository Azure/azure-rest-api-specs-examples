
/**
 * Samples for CassandraDataCenters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBManagedCassandraDataCenterGet.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraDataCenterGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraDataCenters()
            .getWithResponse("cassandra-prod-rg", "cassandra-prod", "dc1", com.azure.core.util.Context.NONE);
    }
}
