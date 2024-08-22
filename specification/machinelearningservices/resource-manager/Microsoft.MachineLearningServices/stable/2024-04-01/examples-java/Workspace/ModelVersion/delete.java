
/**
 * Samples for ModelVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ModelVersion/delete.json
     */
    /**
     * Sample code: Delete Workspace Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
