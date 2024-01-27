
/**
 * Samples for TableResources GetTable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/CosmosDBTableGet.json
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
