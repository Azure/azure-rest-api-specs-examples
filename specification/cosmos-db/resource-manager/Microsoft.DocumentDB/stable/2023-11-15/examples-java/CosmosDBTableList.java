
/**
 * Samples for TableResources ListTables.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/CosmosDBTableList.json
     */
    /**
     * Sample code: CosmosDBTableList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBTableList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getTableResources().listTables("rgName", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
