
/**
 * Samples for DefenderForAISettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetDefenderForAISetting.json
     */
    /**
     * Sample code: GetDefenderForAISetting.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getDefenderForAISetting(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.defenderForAISettings().getWithResponse("resourceGroupName", "accountName", "Default",
            com.azure.core.util.Context.NONE);
    }
}
