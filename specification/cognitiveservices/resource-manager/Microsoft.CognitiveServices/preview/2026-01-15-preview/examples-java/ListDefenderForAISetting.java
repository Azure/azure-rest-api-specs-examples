
/**
 * Samples for DefenderForAISettings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ListDefenderForAISetting.json
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
