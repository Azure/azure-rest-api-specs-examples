
/**
 * Samples for DatabaseAccounts ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountGetUsages.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountGetUsages.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountGetUsages(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listUsages("rg1", "ddb1", "name.value eq 'Storage'",
            com.azure.core.util.Context.NONE);
    }
}
