
/**
 * Samples for FeaturesetContainers GetEntity.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetContainer/getEntity.json
     */
    /**
     * Sample code: GetEntity Workspace Featureset Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getEntityWorkspaceFeaturesetContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetContainers().getEntityWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
