import com.azure.resourcemanager.machinelearning.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.machinelearning.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/PrivateEndpointConnection/createOrUpdate.json
     */
    /**
     * Sample code: WorkspacePutPrivateEndpointConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void workspacePutPrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .privateEndpointConnections()
            .define("{privateEndpointConnectionName}")
            .withExistingWorkspace("rg-1234", "testworkspace")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved"))
            .create();
    }
}
