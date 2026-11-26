
/**
 * Samples for Workspaces ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/listKeys.json
     */
    /**
     * Sample code: List Workspace Keys.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listWorkspaceKeys(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().listKeysWithResponse("testrg123", "workspaces123", com.azure.core.util.Context.NONE);
    }
}
