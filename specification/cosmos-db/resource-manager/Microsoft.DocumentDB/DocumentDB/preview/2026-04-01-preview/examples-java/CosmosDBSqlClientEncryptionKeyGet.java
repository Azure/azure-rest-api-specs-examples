
/**
 * Samples for SqlResources GetClientEncryptionKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlClientEncryptionKeyGet.json
     */
    /**
     * Sample code: CosmosDBClientEncryptionKeyGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBClientEncryptionKeyGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getClientEncryptionKeyWithResponse("rgName", "accountName",
            "databaseName", "cekName", com.azure.core.util.Context.NONE);
    }
}
