
/**
 * Samples for FeaturestoreEntityContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturestoreEntityContainer/delete.json
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
