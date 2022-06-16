import com.azure.core.util.Context;

/** Samples for CassandraResources MigrateCassandraTableToManualThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCassandraTableMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBCassandraTableMigrateToManualThroughput.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraTableMigrateToManualThroughput(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .migrateCassandraTableToManualThroughput("rg1", "ddb1", "keyspaceName", "tableName", Context.NONE);
    }
}
