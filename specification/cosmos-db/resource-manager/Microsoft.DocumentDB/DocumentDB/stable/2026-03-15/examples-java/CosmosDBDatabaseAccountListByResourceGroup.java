
/**
 * Samples for DatabaseAccounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListByResourceGroup.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBDatabaseAccountListByResourceGroup(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
