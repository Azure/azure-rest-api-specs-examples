
/**
 * Samples for CassandraResources MigrateCassandraTableToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBCassandraTableMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBCassandraTableMigrateToAutoscale.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableMigrateToAutoscale(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraResources().migrateCassandraTableToAutoscale(
            "rg1", "ddb1", "keyspaceName", "tableName", com.azure.core.util.Context.NONE);
    }
}
