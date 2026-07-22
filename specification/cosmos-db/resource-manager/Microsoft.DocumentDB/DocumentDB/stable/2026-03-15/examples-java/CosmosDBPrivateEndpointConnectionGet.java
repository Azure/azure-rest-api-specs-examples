
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBPrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("rg1", "ddb1",
            "privateEndpointConnectionName", com.azure.core.util.Context.NONE);
    }
}
