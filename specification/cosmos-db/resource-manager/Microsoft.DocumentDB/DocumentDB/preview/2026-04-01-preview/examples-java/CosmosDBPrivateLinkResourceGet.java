
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBPrivateLinkResourceGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getPrivateLinkResources().getWithResponse("rg1", "ddb1", "sql",
            com.azure.core.util.Context.NONE);
    }
}
