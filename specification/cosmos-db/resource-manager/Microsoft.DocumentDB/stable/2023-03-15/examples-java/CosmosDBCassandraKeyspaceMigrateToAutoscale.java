/** Samples for CassandraResources MigrateCassandraKeyspaceToAutoscale. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBCassandraKeyspaceMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceMigrateToAutoscale.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCassandraKeyspaceMigrateToAutoscale(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraResources()
            .migrateCassandraKeyspaceToAutoscale("rg1", "ddb1", "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
