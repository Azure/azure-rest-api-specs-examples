
import com.azure.resourcemanager.durabletask.models.OptionalPropertiesUpdateableProperties;
import com.azure.resourcemanager.durabletask.models.PrivateEndpointConnection;
import com.azure.resourcemanager.durabletask.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.durabletask.models.PrivateLinkServiceConnectionState;

/**
 * Samples for Schedulers UpdatePrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PrivateEndpointConnections_Update.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Update.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        privateEndpointConnectionsUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        PrivateEndpointConnection resource
            = manager.schedulers().getPrivateEndpointConnectionWithResponse("rgdurabletask", "testscheduler",
                "spzckqrbhfnabu", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new OptionalPropertiesUpdateableProperties().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)))
            .apply();
    }
}
