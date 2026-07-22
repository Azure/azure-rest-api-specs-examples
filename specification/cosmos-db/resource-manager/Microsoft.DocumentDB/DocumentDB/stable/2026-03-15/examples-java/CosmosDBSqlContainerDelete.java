
/**
 * Samples for SqlResources DeleteSqlContainer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerDelete.json
     */
    /**
     * Sample code: CosmosDBSqlContainerDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().deleteSqlContainer("rg1", "ddb1", "databaseName", "containerName",
            com.azure.core.util.Context.NONE);
    }
}
