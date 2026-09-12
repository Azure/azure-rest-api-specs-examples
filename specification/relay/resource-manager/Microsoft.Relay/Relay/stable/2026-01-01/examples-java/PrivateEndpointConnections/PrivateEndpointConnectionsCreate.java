
import com.azure.resourcemanager.relay.models.ConnectionState;
import com.azure.resourcemanager.relay.models.PrivateEndpoint;
import com.azure.resourcemanager.relay.models.PrivateLinkConnectionStatus;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateEndpointConnections/PrivateEndpointConnectionsCreate.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionCreate.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void nameSpacePrivateEndPointConnectionCreate(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.privateEndpointConnections().define("{privateEndpointConnection name}")
            .withExistingNamespace("resourcegroup", "example-RelayNamespace-5849")
            .withPrivateEndpoint(new PrivateEndpoint().withId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Network/privateEndpoints/ali-relay-pve-1"))
            .withPrivateLinkServiceConnectionState(
                new ConnectionState().withStatus(PrivateLinkConnectionStatus.APPROVED).withDescription("You may pass"))
            .create();
    }
}
