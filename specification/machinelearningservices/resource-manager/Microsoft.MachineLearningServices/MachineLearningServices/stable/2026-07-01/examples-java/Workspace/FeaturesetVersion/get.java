
/**
 * Samples for FeaturesetVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturesetVersion/get.json
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
