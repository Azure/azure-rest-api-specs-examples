
/**
 * Samples for DatabaseAccounts ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountGetUsages.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountGetUsages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountGetUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().listUsages("rg1", "ddb1",
            "$filter=name.value eq 'Storage'", com.azure.core.util.Context.NONE);
    }
}
