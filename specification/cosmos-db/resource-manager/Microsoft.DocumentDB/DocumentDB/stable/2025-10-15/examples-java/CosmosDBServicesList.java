
/**
 * Samples for Service List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBServicesList.json
     */
    /**
     * Sample code: CosmosDBServicesList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBServicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().list("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
