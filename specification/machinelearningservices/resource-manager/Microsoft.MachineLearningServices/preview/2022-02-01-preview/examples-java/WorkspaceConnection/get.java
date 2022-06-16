import com.azure.core.util.Context;

/** Samples for WorkspaceConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/get.json
     */
    /**
     * Sample code: GetWorkspaceConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getWorkspaceConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaceConnections().getWithResponse("resourceGroup-1", "workspace-1", "connection-1", Context.NONE);
    }
}
