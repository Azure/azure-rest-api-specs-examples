
/**
 * Samples for SoftDeletedSqlContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlContainerList.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlContainerList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlContainerList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlContainers().listWithResponse("rg1", "West US",
            "softdeleted-cosmosdb-1", "MyDatabase", com.azure.core.util.Context.NONE);
    }
}
