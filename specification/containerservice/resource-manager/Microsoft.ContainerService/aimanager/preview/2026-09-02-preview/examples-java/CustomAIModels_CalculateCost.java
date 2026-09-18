
/**
 * Samples for CustomAIModels CalculateCost.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/CustomAIModels_CalculateCost.json
     */
    /**
     * Sample code: CustomAIModels_CalculateCost.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void customAIModelsCalculateCost(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.customAIModels().calculateCostWithResponse("rg1", "aimanager1", "custom-model1",
            com.azure.core.util.Context.NONE);
    }
}
