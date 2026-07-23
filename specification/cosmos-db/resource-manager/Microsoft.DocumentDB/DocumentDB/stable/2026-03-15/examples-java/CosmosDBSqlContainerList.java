
/**
 * Samples for SqlResources ListSqlContainers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlContainerList.json
     */
    /**
     * Sample code: CosmosDBSqlContainerList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlContainerList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlContainers("rgName", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
