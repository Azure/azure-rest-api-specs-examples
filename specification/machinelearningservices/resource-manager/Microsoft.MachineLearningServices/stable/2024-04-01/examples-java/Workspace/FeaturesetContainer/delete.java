
/**
 * Samples for FeaturesetContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetContainer/delete.json
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
