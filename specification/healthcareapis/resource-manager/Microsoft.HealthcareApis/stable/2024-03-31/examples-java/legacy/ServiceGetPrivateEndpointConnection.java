
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_GetConnection.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        privateEndpointConnectionGetConnection(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateEndpointConnections().getWithResponse("rgname", "service1", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
