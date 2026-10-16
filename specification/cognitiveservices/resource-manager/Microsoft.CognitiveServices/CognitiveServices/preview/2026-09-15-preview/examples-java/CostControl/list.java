
/**
 * Samples for CostControls List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CostControl/list.json
     */
    /**
     * Sample code: List cost controls with all settings.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listCostControlsWithAllSettings(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.costControls().list("foundry-resource-group", "foundry-account", com.azure.core.util.Context.NONE);
    }
}
