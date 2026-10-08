
/**
 * Samples for SoftDeletedSqlContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlContainerGet.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlContainerGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlContainerGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlContainers().getWithResponse("rg1", "West US",
            "softdeleted-cosmosdb-1", "MyDatabase", "MyContainer", com.azure.core.util.Context.NONE);
    }
}
