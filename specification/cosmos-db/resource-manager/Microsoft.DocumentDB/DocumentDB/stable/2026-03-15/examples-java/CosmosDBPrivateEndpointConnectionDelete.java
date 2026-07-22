
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBPrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        deletesAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().delete("rg1", "ddb1", "privateEndpointConnectionName",
            com.azure.core.util.Context.NONE);
    }
}
