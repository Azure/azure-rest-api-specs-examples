import com.azure.core.util.Context;

/** Samples for SqlResources MigrateSqlDatabaseToAutoscale. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlDatabaseMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseMigrateToAutoscale.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlDatabaseMigrateToAutoscale(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .migrateSqlDatabaseToAutoscale("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
