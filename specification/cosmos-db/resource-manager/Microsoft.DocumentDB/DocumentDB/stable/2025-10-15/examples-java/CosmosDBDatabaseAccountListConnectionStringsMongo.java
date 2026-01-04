
/**
 * Samples for DatabaseAccounts ListConnectionStrings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountListConnectionStringsMongo.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListConnectionStringsMongo.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBDatabaseAccountListConnectionStringsMongo(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts()
            .listConnectionStringsWithResponse("rg1", "mongo-ddb1", com.azure.core.util.Context.NONE);
    }
}
