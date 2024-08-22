
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/delete.json
     */
    /**
     * Sample code: Delete Workspace.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().delete("workspace-1234", "testworkspace", null, com.azure.core.util.Context.NONE);
    }
}
