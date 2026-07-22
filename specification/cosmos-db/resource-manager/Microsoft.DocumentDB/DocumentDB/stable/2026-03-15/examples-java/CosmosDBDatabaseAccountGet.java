
/**
 * Samples for DatabaseAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountGet.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().getByResourceGroupWithResponse("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
