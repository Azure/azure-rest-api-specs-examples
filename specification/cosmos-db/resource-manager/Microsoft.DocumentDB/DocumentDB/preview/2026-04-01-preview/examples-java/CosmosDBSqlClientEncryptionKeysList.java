
/**
 * Samples for SqlResources ListClientEncryptionKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlClientEncryptionKeysList.json
     */
    /**
     * Sample code: CosmosDBClientEncryptionKeysList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBClientEncryptionKeysList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listClientEncryptionKeys("rgName", "accountName", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
