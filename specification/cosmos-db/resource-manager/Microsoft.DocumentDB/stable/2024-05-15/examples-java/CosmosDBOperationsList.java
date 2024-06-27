
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/CosmosDBOperationsList.
     * json
     */
    /**
     * Sample code: CosmosDBOperationsList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBOperationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
