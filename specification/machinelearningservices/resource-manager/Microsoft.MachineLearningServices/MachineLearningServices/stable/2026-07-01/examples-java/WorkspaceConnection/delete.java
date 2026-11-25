
/**
 * Samples for WorkspaceConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/WorkspaceConnection/delete.json
     */
    /**
     * Sample code: DeleteWorkspaceConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceConnection(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaceConnections().deleteWithResponse("resourceGroup-1", "workspace-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}
