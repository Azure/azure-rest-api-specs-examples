
import com.azure.resourcemanager.containerserviceaimanager.models.CalculateCostRequest;

/**
 * Samples for AIModels CalculateCost.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIModels_CalculateCost.json
     */
    /**
     * Sample code: AIModels_CalculateCost_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void aIModelsCalculateCostMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIModels().calculateCostWithResponse("eastus", "9806f0c862fdd920", new CalculateCostRequest(),
            com.azure.core.util.Context.NONE);
    }
}
