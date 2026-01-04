
/**
 * Samples for RestorableTables List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBRestorableTableList.json
     */
    /**
     * Sample code: CosmosDBRestorableTableList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableTableList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getRestorableTables().list("WestUS",
            "98a570f2-63db-4117-91f0-366327b7b353", null, null, com.azure.core.util.Context.NONE);
    }
}
