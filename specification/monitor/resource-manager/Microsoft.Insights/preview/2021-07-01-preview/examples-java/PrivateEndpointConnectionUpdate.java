
import com.azure.resourcemanager.monitor.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.monitor.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.monitor.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        approveOrRejectAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateEndpointConnections().createOrUpdate(
            "MyResourceGroup", "MyPrivateLinkScope", "private-endpoint-connection-name",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com")),
            com.azure.core.util.Context.NONE);
    }
}
