
/**
 * Samples for CodeVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/CodeVersion/delete.json
     */
    /**
     * Sample code: Delete Workspace Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
