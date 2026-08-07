
/**
 * Samples for AIModels Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIModels_Get.json
     */
    /**
     * Sample code: AIModels_Get_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void aIModelsGetMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIModels().getWithResponse("eastus", "9806f0c862fdd920", com.azure.core.util.Context.NONE);
    }
}
