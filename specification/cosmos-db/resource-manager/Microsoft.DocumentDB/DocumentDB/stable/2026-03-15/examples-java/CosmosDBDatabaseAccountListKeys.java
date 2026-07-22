
/**
 * Samples for DatabaseAccounts ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountListKeys.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListKeys.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountListKeys(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listKeysWithResponse("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
