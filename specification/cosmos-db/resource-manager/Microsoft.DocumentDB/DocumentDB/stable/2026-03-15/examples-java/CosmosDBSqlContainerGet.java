
/**
 * Samples for SqlResources GetSqlContainer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerGet.json
     */
    /**
     * Sample code: CosmosDBSqlContainerGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlContainerWithResponse("rgName", "ddb1", "databaseName",
            "containerName", com.azure.core.util.Context.NONE);
    }
}
