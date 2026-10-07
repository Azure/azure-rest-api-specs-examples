
/**
 * Samples for SoftDeletedSqlContainers Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlContainerPurge.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlContainerPurge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlContainerPurge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlContainers().purge("rg1", "West US", "softdeleted-cosmosdb-1",
            "MyDatabase", "MyContainer", null, com.azure.core.util.Context.NONE);
    }
}
