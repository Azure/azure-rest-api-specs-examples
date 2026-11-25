
/**
 * Samples for CodeContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/CodeContainer/list.json
     */
    /**
     * Sample code: List Workspace Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeContainers().list("testrg123", "testworkspace", null, com.azure.core.util.Context.NONE);
    }
}
