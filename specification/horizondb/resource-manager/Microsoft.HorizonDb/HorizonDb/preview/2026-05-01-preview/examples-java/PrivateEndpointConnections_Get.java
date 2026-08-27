
/**
 * Samples for HorizonDbPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: Get a private endpoint connection.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAPrivateEndpointConnection(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbPrivateEndpointConnections().getWithResponse("exampleresourcegroup", "examplecluster",
            "exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e", com.azure.core.util.Context.NONE);
    }
}
