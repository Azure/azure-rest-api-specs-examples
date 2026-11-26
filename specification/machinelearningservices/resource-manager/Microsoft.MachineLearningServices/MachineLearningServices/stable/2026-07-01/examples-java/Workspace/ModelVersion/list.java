
/**
 * Samples for ModelVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ModelVersion/list.json
     */
    /**
     * Sample code: List Workspace Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelVersions().list("test-rg", "my-aml-workspace", "string", null, "string", 1, "string", "string", 1,
            "string", "string", null, null, com.azure.core.util.Context.NONE);
    }
}
