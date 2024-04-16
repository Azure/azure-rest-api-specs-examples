
import com.azure.resourcemanager.healthcareapis.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.healthcareapis.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceCreatePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_CreateOrUpdate.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void privateEndpointConnectionCreateOrUpdate(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateEndpointConnections().define("myConnection").withExistingService("rgname", "service1")
            .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED).withDescription("Auto-Approved"))
            .create();
    }
}
