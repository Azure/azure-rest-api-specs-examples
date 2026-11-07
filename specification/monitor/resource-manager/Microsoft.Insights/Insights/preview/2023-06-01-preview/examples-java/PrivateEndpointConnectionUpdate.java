
import com.azure.resourcemanager.monitor.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.monitor.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.monitor.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void approveOrRejectAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().createOrUpdate("MyResourceGroup", "MyPrivateLinkScope",
            "private-endpoint-connection-name",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com")),
            com.azure.core.util.Context.NONE);
    }
}
