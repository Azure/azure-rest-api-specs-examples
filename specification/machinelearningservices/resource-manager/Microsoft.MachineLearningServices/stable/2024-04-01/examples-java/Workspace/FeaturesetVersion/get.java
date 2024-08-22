
/**
 * Samples for FeaturesetVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetVersion/get.json
     */
    /**
     * Sample code: Get Workspace Featureset Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceFeaturesetVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
