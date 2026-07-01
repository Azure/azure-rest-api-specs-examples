
/**
 * Samples for RaiExternalSafetyProvider Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetRaiExternalSafetyProvider.json
     */
    /**
     * Sample code: GetRaiExternalSafetyProvider.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getRaiExternalSafetyProvider(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiExternalSafetyProviders().getWithResponse("safetyProviderName", com.azure.core.util.Context.NONE);
    }
}
