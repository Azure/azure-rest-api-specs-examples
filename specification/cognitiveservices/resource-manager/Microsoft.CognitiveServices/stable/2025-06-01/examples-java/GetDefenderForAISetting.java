
/**
 * Samples for DefenderForAISettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * GetDefenderForAISetting.json
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
