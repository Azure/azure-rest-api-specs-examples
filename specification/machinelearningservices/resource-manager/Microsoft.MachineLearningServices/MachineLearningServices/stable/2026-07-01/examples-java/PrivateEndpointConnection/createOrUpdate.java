
import com.azure.resourcemanager.machinelearning.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.machinelearning.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/PrivateEndpointConnection/createOrUpdate.json
     */
    /**
     * Sample code: WorkspacePutPrivateEndpointConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void workspacePutPrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.privateEndpointConnections().define("{privateEndpointConnectionName}")
            .withExistingWorkspace("rg-1234", "testworkspace")
            .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                .withDescription("Auto-Approved").withStatus(PrivateEndpointServiceConnectionStatus.APPROVED))
            .create();
    }
}
