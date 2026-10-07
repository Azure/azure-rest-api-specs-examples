
/**
 * Samples for SoftDeletedDatabaseAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedDatabaseAccountGet.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedDatabaseAccountGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedDatabaseAccountGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedDatabaseAccounts().getWithResponse("rg1", "West US",
            "softdeleted-cosmosdb-1", com.azure.core.util.Context.NONE);
    }
}
