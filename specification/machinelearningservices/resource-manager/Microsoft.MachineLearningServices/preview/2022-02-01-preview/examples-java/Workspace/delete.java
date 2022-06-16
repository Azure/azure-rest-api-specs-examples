import com.azure.core.util.Context;

/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/delete.json
     */
    /**
     * Sample code: Delete Workspace.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteWorkspace(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().delete("workspace-1234", "testworkspace", Context.NONE);
    }
}
