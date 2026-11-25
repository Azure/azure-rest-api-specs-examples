
/**
 * Samples for FeaturesetVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturesetVersion/delete.json
     */
    /**
     * Sample code: Delete Workspace Featureset Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceFeaturesetVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetVersions().delete("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
