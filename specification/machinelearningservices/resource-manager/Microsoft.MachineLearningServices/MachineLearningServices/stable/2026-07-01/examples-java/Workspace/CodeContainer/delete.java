
/**
 * Samples for CodeContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/CodeContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeContainers().deleteWithResponse("testrg123", "testworkspace", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}
