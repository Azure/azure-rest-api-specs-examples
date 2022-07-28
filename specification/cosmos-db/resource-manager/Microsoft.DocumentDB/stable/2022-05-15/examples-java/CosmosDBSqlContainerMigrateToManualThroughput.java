import com.azure.core.util.Context;

/** Samples for SqlResources MigrateSqlContainerToManualThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBSqlContainerMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBSqlContainerMigrateToManualThroughput.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerMigrateToManualThroughput(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .migrateSqlContainerToManualThroughput("rg1", "ddb1", "databaseName", "containerName", Context.NONE);
    }
}
