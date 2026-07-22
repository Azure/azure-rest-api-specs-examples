
/**
 * Samples for SqlResources GetSqlContainerThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerThroughputGet.json
     */
    /**
     * Sample code: CosmosDBSqlContainerThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlContainerThroughputWithResponse("rg1", "ddb1", "databaseName",
            "containerName", com.azure.core.util.Context.NONE);
    }
}
