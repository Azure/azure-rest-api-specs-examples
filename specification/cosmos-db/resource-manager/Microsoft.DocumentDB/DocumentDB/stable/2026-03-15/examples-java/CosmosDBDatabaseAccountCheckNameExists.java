
/**
 * Samples for DatabaseAccounts CheckNameExists.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountCheckNameExists.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountCheckNameExists.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountCheckNameExists(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().checkNameExistsWithResponse("ddb1",
            com.azure.core.util.Context.NONE);
    }
}
