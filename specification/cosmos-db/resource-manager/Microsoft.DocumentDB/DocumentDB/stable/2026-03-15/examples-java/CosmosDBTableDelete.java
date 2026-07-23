
/**
 * Samples for TableResources DeleteTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBTableDelete.json
     */
    /**
     * Sample code: CosmosDBTableDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().deleteTable("rg1", "ddb1", "tableName",
            com.azure.core.util.Context.NONE);
    }
}
