import com.azure.core.util.Context;

/** Samples for WorkspaceConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/list.json
     */
    /**
     * Sample code: ListWorkspaceConnections.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listWorkspaceConnections(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaceConnections().list("resourceGroup-1", "workspace-1", "www.facebook.com", "ACR", Context.NONE);
    }
}
