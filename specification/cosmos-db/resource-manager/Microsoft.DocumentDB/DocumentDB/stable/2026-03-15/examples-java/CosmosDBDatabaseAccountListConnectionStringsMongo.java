
/**
 * Samples for DatabaseAccounts ListConnectionStrings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountListConnectionStringsMongo.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListConnectionStringsMongo.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBDatabaseAccountListConnectionStringsMongo(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listConnectionStringsWithResponse("rg1", "mongo-ddb1",
            com.azure.core.util.Context.NONE);
    }
}
