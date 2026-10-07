
/**
 * Samples for SoftDeletedSqlContainers Restore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlContainerRestore.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlContainerRestore.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlContainerRestore(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlContainers().restore("rg1", "West US", "softdeleted-cosmosdb-1",
            "MyDatabase", "MyContainer", null, com.azure.core.util.Context.NONE);
    }
}
