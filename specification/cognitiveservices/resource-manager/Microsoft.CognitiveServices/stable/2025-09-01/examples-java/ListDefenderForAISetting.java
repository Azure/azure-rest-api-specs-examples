
/**
 * Samples for DefenderForAISettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListDefenderForAISetting.json
     */
    /**
     * Sample code: ListDefenderForAISetting.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listDefenderForAISetting(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.defenderForAISettings().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
