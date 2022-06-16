import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_GetConnection.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void privateEndpointConnectionGetConnection(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateEndpointConnections().getWithResponse("rgname", "service1", "myConnection", Context.NONE);
    }
}
