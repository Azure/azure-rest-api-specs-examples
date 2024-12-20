
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        privateEndpointConnectionsDelete(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateEndpointConnections().delete("rgname", "service1", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
