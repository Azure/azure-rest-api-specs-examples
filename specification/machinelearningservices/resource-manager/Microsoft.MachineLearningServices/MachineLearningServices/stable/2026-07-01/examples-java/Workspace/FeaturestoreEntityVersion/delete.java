
/**
 * Samples for FeaturestoreEntityVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturestoreEntityVersion/delete.json
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
