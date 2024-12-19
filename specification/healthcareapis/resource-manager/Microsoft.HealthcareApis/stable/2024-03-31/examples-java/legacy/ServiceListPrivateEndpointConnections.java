
/**
 * Samples for PrivateEndpointConnections ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceListPrivateEndpointConnections.json
     */
    /**
     * Sample code: PrivateEndpointConnection_List.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        privateEndpointConnectionList(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateEndpointConnections().listByService("rgname", "service1", com.azure.core.util.Context.NONE);
    }
}
