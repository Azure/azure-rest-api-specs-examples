
/**
 * Samples for ModelVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ModelVersion/list.json
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
