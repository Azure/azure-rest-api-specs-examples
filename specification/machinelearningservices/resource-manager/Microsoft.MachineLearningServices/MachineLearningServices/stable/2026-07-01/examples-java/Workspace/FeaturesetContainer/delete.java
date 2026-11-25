
/**
 * Samples for FeaturesetContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturesetContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Featureset Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceFeaturesetContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetContainers().delete("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
