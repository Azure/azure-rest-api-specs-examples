
/**
 * Samples for SoftDeletedDatabaseAccounts Restore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedDatabaseAccountRestore.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedDatabaseAccountRestore.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSoftDeletedDatabaseAccountRestore(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedDatabaseAccounts().restore("rg2", "West US", "softdeleted-cosmosdb-2",
            null, com.azure.core.util.Context.NONE);
    }
}
