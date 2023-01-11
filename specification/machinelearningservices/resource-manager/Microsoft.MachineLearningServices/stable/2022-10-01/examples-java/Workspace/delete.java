import com.azure.core.util.Context;

/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Workspace/delete.json
     */
    /**
     * Sample code: Delete Workspace.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().delete("workspace-1234", "testworkspace", Context.NONE);
    }
}
