/** Samples for PrivateEndpointConnections ListByDatabaseAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBPrivateEndpointConnectionListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .listByDatabaseAccount("rg1", "ddb1", com.azure.core.util.Context.NONE);
    }
}
