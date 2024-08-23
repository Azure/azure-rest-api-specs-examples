
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/listByResourceGroup.json
     */
    /**
     * Sample code: Get Workspaces by Resource Group.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspacesByResourceGroup(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().listByResourceGroup("workspace-1234", null, com.azure.core.util.Context.NONE);
    }
}
