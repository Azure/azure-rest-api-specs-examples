
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/get.json
     */
    /**
     * Sample code: Get Workspace.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("workspace-1234", "testworkspace",
            com.azure.core.util.Context.NONE);
    }
}
