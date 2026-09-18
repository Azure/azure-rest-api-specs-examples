
/**
 * Samples for CustomAIModels Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/CustomAIModels_Get.json
     */
    /**
     * Sample code: Get a CustomAIModel.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void getACustomAIModel(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.customAIModels().getWithResponse("rg1", "aimanager1", "custom-model1",
            com.azure.core.util.Context.NONE);
    }
}
