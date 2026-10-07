
/**
 * Samples for SqlResources ListSqlContainers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSqlContainerList.json
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
