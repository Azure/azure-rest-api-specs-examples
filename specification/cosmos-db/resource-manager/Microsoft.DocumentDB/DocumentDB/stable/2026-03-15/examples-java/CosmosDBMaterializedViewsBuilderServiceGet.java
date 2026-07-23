
/**
 * Samples for Service Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMaterializedViewsBuilderServiceGet.json
     */
    /**
     * Sample code: MaterializedViewsBuilderServiceGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void materializedViewsBuilderServiceGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().getWithResponse("rg1", "ddb1", "MaterializedViewsBuilder",
            com.azure.core.util.Context.NONE);
    }
}
