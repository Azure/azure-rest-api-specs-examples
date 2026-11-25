
/**
 * Samples for Features List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Feature/list.json
     */
    /**
     * Sample code: List Feature.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listFeature(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.features().list("test-rg", "my-aml-workspace", "string", "string", null, "string", "string", "string",
            null, null, com.azure.core.util.Context.NONE);
    }
}
