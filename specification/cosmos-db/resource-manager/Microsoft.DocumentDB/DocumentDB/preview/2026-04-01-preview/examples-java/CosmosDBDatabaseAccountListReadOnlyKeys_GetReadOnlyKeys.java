
/**
 * Samples for DatabaseAccounts GetReadOnlyKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBDatabaseAccountListReadOnlyKeys_GetReadOnlyKeys.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListReadOnlyKeys.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountListReadOnlyKeys(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().getReadOnlyKeysWithResponse("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
