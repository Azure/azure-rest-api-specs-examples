
/**
 * Samples for CostControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CostControl/getLegacy.json
     */
    /**
     * Sample code: Read a legacy cost control without authoring its legacy values.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void readALegacyCostControlWithoutAuthoringItsLegacyValues(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.costControls().getWithResponse("foundry-resource-group", "foundry-account", "legacy-budget",
            com.azure.core.util.Context.NONE);
    }
}
