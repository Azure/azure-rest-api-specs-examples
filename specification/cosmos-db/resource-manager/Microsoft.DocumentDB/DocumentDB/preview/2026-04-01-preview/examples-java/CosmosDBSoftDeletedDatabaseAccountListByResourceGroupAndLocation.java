
/**
 * Samples for SoftDeletedDatabaseAccounts ListByResourceGroupAndLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedDatabaseAccountListByResourceGroupAndLocation.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedDatabaseAccountListByResourceGroupAndLocation.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedDatabaseAccountListByResourceGroupAndLocation(
        com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedDatabaseAccounts().listByResourceGroupAndLocationWithResponse("rg1",
            "West US", com.azure.core.util.Context.NONE);
    }
}
