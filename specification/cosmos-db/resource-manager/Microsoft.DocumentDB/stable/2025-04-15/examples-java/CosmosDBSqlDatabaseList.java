
/**
 * Samples for SqlResources ListSqlDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBSqlDatabaseList.
     * json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlDatabaseList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().listSqlDatabases("rgName", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
