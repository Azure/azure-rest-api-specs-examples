
/**
 * Samples for HorizonDbPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: Delete a private endpoint connection.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAPrivateEndpointConnection(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbPrivateEndpointConnections().delete("exampleresourcegroup", "examplecluster",
            "exampleprivateendpointconnection.1fa229cd-bf3f-47f0-8c49-afb36723997e", com.azure.core.util.Context.NONE);
    }
}
