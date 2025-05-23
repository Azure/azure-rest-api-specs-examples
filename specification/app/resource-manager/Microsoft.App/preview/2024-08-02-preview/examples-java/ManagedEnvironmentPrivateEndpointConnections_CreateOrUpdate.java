
import com.azure.resourcemanager.appcontainers.models.PrivateEndpointConnection;
import com.azure.resourcemanager.appcontainers.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.appcontainers.models.PrivateLinkServiceConnectionState;

/**
 * Samples for ManagedEnvironmentPrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/
     * ManagedEnvironmentPrivateEndpointConnections_CreateOrUpdate.json
     */
    /**
     * Sample code: Update a Private Endpoint Connection by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void updateAPrivateEndpointConnectionByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        PrivateEndpointConnection resource = manager.managedEnvironmentPrivateEndpointConnections()
            .getWithResponse("examplerg", "managedEnv", "jlaw-demo1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED).withActionsRequired("None")).apply();
    }
}
