
/**
 * Samples for Workspaces ResyncKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/resyncKeys.json
     */
    /**
     * Sample code: Resync Workspace Keys.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void resyncWorkspaceKeys(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().resyncKeys("testrg123", "workspaces123", com.azure.core.util.Context.NONE);
    }
}
