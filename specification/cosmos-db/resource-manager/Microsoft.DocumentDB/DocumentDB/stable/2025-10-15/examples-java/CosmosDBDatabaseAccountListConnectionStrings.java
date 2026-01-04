
/**
 * Samples for DatabaseAccounts ListConnectionStrings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountListConnectionStrings.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListConnectionStrings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBDatabaseAccountListConnectionStrings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts()
            .listConnectionStringsWithResponse("rg1", "ddb1", com.azure.core.util.Context.NONE);
    }
}
