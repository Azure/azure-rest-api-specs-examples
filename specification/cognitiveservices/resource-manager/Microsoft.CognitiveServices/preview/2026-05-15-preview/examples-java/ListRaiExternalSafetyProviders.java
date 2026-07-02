
/**
 * Samples for RaiExternalSafetyProvidersOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListRaiExternalSafetyProviders.json
     */
    /**
     * Sample code: ListRaiExternalSafetyProviders.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listRaiExternalSafetyProviders(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiExternalSafetyProvidersOperations().list(com.azure.core.util.Context.NONE);
    }
}
