
/**
 * Samples for WorkspaceFeatures List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/WorkspaceFeature/list.json
     */
    /**
     * Sample code: List Workspace features.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listWorkspaceFeatures(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaceFeatures().list("myResourceGroup", "testworkspace", com.azure.core.util.Context.NONE);
    }
}
