
/**
 * Samples for SqlResources GetClientEncryptionKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBSqlClientEncryptionKeyGet.json
     */
    /**
     * Sample code: CosmosDBClientEncryptionKeyGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBClientEncryptionKeyGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getClientEncryptionKeyWithResponse(
            "rgName", "accountName", "databaseName", "cekName", com.azure.core.util.Context.NONE);
    }
}
