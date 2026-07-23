
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMaterializedViewsBuilderServiceDelete.json
     */
    /**
     * Sample code: MaterializedViewsBuilderServiceDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void materializedViewsBuilderServiceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().delete("rg1", "ddb1", "MaterializedViewsBuilder",
            com.azure.core.util.Context.NONE);
    }
}
