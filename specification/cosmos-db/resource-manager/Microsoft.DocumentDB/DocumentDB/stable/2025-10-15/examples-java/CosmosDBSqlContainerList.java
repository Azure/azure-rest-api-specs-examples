
/**
 * Samples for SqlResources ListSqlContainers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlContainerList.json
     */
    /**
     * Sample code: CosmosDBSqlContainerList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().listSqlContainers("rgName", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
