
/**
 * Samples for DatabaseAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountDelete.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().delete("rg1", "ddb1", com.azure.core.util.Context.NONE);
    }
}
