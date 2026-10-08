
/**
 * Samples for SoftDeletedDatabaseAccounts ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedDatabaseAccountListByLocation.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedDatabaseAccountListByLocation.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSoftDeletedDatabaseAccountListByLocation(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedDatabaseAccounts().listByLocationWithResponse("West US",
            com.azure.core.util.Context.NONE);
    }
}
