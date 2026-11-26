
/**
 * Samples for FeaturestoreEntityVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturestoreEntityVersion/get.json
     */
    /**
     * Sample code: Get Workspace Featurestore Entity Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getWorkspaceFeaturestoreEntityVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
