
/**
 * Samples for DatabaseAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountList.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().list(com.azure.core.util.Context.NONE);
    }
}
