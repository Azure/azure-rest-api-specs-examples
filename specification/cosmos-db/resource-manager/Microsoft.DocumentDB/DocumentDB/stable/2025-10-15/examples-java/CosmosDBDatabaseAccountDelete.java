
/**
 * Samples for DatabaseAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountDelete.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().delete("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
