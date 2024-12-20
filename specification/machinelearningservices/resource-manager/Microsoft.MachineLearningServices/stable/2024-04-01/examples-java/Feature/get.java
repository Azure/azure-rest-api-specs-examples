
/**
 * Samples for Features Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Feature/get.json
     */
    /**
     * Sample code: Get Feature.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getFeature(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.features().getWithResponse("test-rg", "my-aml-workspace", "string", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
