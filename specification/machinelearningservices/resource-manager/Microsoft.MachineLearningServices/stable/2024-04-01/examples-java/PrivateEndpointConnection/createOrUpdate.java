
import com.azure.resourcemanager.machinelearning.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.machinelearning.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/PrivateEndpointConnection/createOrUpdate.json
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
                .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED).withDescription("Auto-Approved"))
            .create();
    }
}
