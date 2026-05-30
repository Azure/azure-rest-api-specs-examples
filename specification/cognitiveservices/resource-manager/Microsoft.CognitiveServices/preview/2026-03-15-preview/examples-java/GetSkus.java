
/**
 * Samples for ResourceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetSkus.json
     */
    /**
     * Sample code: Regenerate Keys.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void regenerateKeys(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.resourceSkus().list(com.azure.core.util.Context.NONE);
    }
}
