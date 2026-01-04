
/**
 * Samples for Database ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseGetUsages.json
     */
    /**
     * Sample code: CosmosDBDatabaseGetUsages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseGetUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabases().listUsages("rg1", "ddb1", "databaseRid",
            "$filter=name.value eq 'Storage'", com.azure.core.util.Context.NONE);
    }
}
