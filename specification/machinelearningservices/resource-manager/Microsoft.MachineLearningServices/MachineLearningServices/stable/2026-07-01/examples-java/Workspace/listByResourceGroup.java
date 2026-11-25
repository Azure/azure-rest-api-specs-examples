
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/listByResourceGroup.json
     */
    /**
     * Sample code: Get Workspaces by Resource Group.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspacesByResourceGroup(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().listByResourceGroup("workspace-1234", null, null, null, com.azure.core.util.Context.NONE);
    }
}
