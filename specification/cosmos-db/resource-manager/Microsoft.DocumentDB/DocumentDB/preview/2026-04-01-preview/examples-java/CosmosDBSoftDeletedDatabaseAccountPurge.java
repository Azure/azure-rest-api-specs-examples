
/**
 * Samples for SoftDeletedDatabaseAccounts Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedDatabaseAccountPurge.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedDatabaseAccountPurge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedDatabaseAccountPurge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedDatabaseAccounts().purge("rg1", "West US", "softdeleted-cosmosdb-1", null,
            com.azure.core.util.Context.NONE);
    }
}
