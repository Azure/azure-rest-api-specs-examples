import com.azure.core.util.Context;

/** Samples for WorkspaceConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/delete.json
     */
    /**
     * Sample code: DeleteWorkspaceConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteWorkspaceConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .workspaceConnections()
            .deleteWithResponse("resourceGroup-1", "workspace-1", "connection-1", Context.NONE);
    }
}
