
/**
 * Samples for FeaturestoreEntityVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturestoreEntityVersion/delete.json
     */
    /**
     * Sample code: Delete Workspace Featurestore Entity Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteWorkspaceFeaturestoreEntityVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityVersions().delete("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
