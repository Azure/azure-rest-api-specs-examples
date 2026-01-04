
/**
 * Samples for TableResources GetTable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBTableGet.json
     */
    /**
     * Sample code: CosmosDBTableGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBTableGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getTableResources().getTableWithResponse("rg1", "ddb1",
            "tableName", com.azure.core.util.Context.NONE);
    }
}
