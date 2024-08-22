
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/get.json
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
