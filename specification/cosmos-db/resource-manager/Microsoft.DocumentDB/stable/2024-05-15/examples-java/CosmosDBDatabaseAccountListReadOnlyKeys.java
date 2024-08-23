
/**
 * Samples for DatabaseAccounts GetReadOnlyKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBDatabaseAccountListReadOnlyKeys.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListReadOnlyKeys.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountListReadOnlyKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().getReadOnlyKeysWithResponse("rg1",
            "ddb1", com.azure.core.util.Context.NONE);
    }
}
