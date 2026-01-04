
/**
 * Samples for TableResources MigrateTableToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBTableMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBTableMigrateToManualThroughput.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBTableMigrateToManualThroughput(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getTableResources().migrateTableToManualThroughput("rg1",
            "ddb1", "tableName", com.azure.core.util.Context.NONE);
    }
}
