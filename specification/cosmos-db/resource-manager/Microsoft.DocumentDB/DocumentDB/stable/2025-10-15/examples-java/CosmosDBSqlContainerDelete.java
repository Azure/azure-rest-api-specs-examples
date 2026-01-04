
/**
 * Samples for SqlResources DeleteSqlContainer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlContainerDelete.json
     */
    /**
     * Sample code: CosmosDBSqlContainerDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().deleteSqlContainer("rg1", "ddb1",
            "databaseName", "containerName", com.azure.core.util.Context.NONE);
    }
}
