
/**
 * Samples for CodeVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/CodeVersion/list.json
     */
    /**
     * Sample code: List Workspace Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().list("test-rg", "my-aml-workspace", "string", "string", 1, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
