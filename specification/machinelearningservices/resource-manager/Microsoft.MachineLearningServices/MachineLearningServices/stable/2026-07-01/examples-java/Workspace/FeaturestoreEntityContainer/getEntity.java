
/**
 * Samples for FeaturestoreEntityContainers GetEntity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturestoreEntityContainer/getEntity.json
     */
    /**
     * Sample code: GetEntity Workspace Featurestore Entity Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getEntityWorkspaceFeaturestoreEntityContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityContainers().getEntityWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
