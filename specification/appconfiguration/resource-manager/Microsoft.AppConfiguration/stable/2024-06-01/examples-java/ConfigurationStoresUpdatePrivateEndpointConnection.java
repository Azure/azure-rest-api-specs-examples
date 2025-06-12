
import com.azure.resourcemanager.appconfiguration.models.ConnectionStatus;
import com.azure.resourcemanager.appconfiguration.models.PrivateEndpointConnection;
import com.azure.resourcemanager.appconfiguration.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/
     * ConfigurationStoresUpdatePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Update.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        privateEndpointConnectionUpdate(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        PrivateEndpointConnection resource = manager.privateEndpointConnections()
            .getWithResponse("myResourceGroup", "contoso", "myConnection", com.azure.core.util.Context.NONE).getValue();
        resource.update().withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
            .withStatus(ConnectionStatus.APPROVED).withDescription("Auto-Approved")).apply();
    }
}
