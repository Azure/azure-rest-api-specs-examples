
/**
 * Samples for PrivateLinkResources ListByDatabaseAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBPrivateLinkResourceListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getPrivateLinkResources().listByDatabaseAccount("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
