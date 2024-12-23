
/**
 * Samples for PrivateLinkResources ListByDatabaseAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/
     * CosmosDBPrivateLinkResourceListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getPrivateLinkResources().listByDatabaseAccount("rg1",
            "ddb1", com.azure.core.util.Context.NONE);
    }
}
