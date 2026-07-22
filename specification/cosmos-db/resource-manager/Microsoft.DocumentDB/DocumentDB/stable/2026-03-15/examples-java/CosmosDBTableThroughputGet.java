
/**
 * Samples for TableResources GetTableThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBTableThroughputGet.json
     */
    /**
     * Sample code: CosmosDBTableThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().getTableThroughputWithResponse("rg1", "ddb1", "tableName",
            com.azure.core.util.Context.NONE);
    }
}
