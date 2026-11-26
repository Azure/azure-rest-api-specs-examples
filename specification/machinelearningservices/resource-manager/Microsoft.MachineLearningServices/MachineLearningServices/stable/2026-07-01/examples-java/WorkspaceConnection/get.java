
/**
 * Samples for WorkspaceConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/WorkspaceConnection/get.json
     */
    /**
     * Sample code: GetWorkspaceConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceConnection(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaceConnections().getWithResponse("resourceGroup-1", "workspace-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}
