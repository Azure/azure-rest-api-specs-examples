
/**
 * Samples for SqlResources GetSqlContainerThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlContainerThroughputGet.json
     */
    /**
     * Sample code: CosmosDBSqlContainerThroughputGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlContainerThroughputWithResponse(
            "rg1", "ddb1", "databaseName", "containerName", com.azure.core.util.Context.NONE);
    }
}
