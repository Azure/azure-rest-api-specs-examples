
import com.azure.resourcemanager.appcontainers.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.appcontainers.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.appcontainers.models.PrivateLinkServiceConnectionState;

/**
 * Samples for ContainerAppPrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerAppPrivateEndpointConnections_CreateOrUpdate.json
     */
    /**
     * Sample code: Update a Private Endpoint Connection by Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void updateAPrivateEndpointConnectionByContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppPrivateEndpointConnections().createOrUpdate("examplerg", "testcontainerapp0",
            "test-private-endpoint-connection",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withActionsRequired("None")),
            com.azure.core.util.Context.NONE);
    }
}
