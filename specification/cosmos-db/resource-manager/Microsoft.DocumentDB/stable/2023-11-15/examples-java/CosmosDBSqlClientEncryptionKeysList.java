
/**
 * Samples for SqlResources ListClientEncryptionKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/
     * CosmosDBSqlClientEncryptionKeysList.json
     */
    /**
     * Sample code: CosmosDBClientEncryptionKeysList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBClientEncryptionKeysList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().listClientEncryptionKeys("rgName",
            "accountName", "databaseName", com.azure.core.util.Context.NONE);
    }
}
