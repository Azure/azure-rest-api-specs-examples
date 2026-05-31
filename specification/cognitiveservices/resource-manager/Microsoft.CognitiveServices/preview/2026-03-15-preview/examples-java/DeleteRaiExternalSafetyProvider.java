
/**
 * Samples for RaiExternalSafetyProvider Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteRaiExternalSafetyProvider.json
     */
    /**
     * Sample code: DeleteRaiTopic.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteRaiTopic(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiExternalSafetyProviders().delete("safetyProviderName", com.azure.core.util.Context.NONE);
    }
}
