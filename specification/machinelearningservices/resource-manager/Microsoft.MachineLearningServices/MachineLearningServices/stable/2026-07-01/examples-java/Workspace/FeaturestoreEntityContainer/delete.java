
/**
 * Samples for FeaturestoreEntityContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturestoreEntityContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Featurestore Entity Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteWorkspaceFeaturestoreEntityContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityContainers().delete("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
