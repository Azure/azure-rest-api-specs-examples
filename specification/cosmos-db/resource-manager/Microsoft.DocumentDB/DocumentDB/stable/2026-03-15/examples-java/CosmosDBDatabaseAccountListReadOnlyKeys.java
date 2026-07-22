
/**
 * Samples for DatabaseAccounts ListReadOnlyKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountListReadOnlyKeys.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListReadOnlyKeys.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountListReadOnlyKeys(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listReadOnlyKeysWithResponse("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
