
/**
 * Samples for TableResources GetTableThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBTableThroughputGet.json
     */
    /**
     * Sample code: CosmosDBTableThroughputGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBTableThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getTableResources().getTableThroughputWithResponse("rg1",
            "ddb1", "tableName", com.azure.core.util.Context.NONE);
    }
}
