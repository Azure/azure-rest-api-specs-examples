
import com.azure.resourcemanager.containerregistry.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.containerregistry.models.ConnectionStatus;
import com.azure.resourcemanager.containerregistry.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/PrivateEndpointConnectionCreateOrUpdate.json
     */
    /**
     * Sample code: PrivateEndpointConnectionCreateOrUpdate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void privateEndpointConnectionCreateOrUpdate(
        com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().createOrUpdate("myResourceGroup", "myRegistry",
            "myConnection",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(ConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved")),
            com.azure.core.util.Context.NONE);
    }
}
