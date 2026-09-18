
import com.azure.resourcemanager.servicenetworking.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.servicenetworking.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.servicenetworking.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.servicenetworking.models.PrivateLinkServiceConnectionStatus;

/**
 * Samples for PrivateEndpointConnectionsInterface Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateEndpointConnectionPut.json
     */
    /**
     * Sample code: Update Private Endpoint Connection.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        updatePrivateEndpointConnection(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.privateEndpointConnectionsInterfaces().update("rg1", "tc1", "pec1",
            new PrivateEndpointConnectionInner().withProperties(new PrivateEndpointConnectionProperties()
                .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateLinkServiceConnectionStatus.APPROVED).withDescription("Approved by admin"))),
            com.azure.core.util.Context.NONE);
    }
}
