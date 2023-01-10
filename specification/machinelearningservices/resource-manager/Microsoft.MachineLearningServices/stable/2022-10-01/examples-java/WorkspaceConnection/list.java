import com.azure.core.util.Context;

/** Samples for WorkspaceConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/WorkspaceConnection/list.json
     */
    /**
     * Sample code: ListWorkspaceConnections.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listWorkspaceConnections(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .workspaceConnections()
            .list("resourceGroup-1", "workspace-1", "www.facebook.com", "ContainerRegistry", Context.NONE);
    }
}
