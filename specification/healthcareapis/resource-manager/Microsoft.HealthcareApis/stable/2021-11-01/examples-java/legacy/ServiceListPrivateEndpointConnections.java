import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceListPrivateEndpointConnections.json
     */
    /**
     * Sample code: PrivateEndpointConnection_List.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void privateEndpointConnectionList(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateEndpointConnections().listByService("rgname", "service1", Context.NONE);
    }
}
