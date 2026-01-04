
/**
 * Samples for SqlResources MigrateSqlContainerToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlContainerMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBSqlContainerMigrateToAutoscale.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerMigrateToAutoscale(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().migrateSqlContainerToAutoscale("rg1",
            "ddb1", "databaseName", "containerName", com.azure.core.util.Context.NONE);
    }
}
