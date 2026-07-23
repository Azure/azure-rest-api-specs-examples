
/**
 * Samples for TableResources GetTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBTableGet.json
     */
    /**
     * Sample code: CosmosDBTableGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().getTableWithResponse("rg1", "ddb1", "tableName",
            com.azure.core.util.Context.NONE);
    }
}
