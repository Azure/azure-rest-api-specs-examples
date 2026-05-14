
import com.azure.resourcemanager.fileshares.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.fileshares.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.fileshares.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/PrivateEndpointConnections_Create.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Create.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        privateEndpointConnectionsCreate(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.privateEndpointConnections().define("privateEndpointConnection1")
            .withExistingFileShare("rgfileshares", "fileshare")
            .withProperties(new PrivateEndpointConnectionProperties()
                .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED).withDescription("Approved by admin")))
            .create();
    }
}
