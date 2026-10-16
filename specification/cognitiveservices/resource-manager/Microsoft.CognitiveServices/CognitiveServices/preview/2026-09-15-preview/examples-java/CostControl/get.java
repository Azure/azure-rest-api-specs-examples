
/**
 * Samples for CostControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CostControl/get.json
     */
    /**
     * Sample code: Get a cost control with all settings.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getACostControlWithAllSettings(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.costControls().getWithResponse("foundry-resource-group", "foundry-account", "production-agents",
            com.azure.core.util.Context.NONE);
    }
}
