
/**
 * Samples for DatabaseAccounts ListConnectionStrings.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountListConnectionStrings.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListConnectionStrings.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBDatabaseAccountListConnectionStrings(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listConnectionStringsWithResponse("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
