
/**
 * Samples for RaiBindings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/ListRaiBindings.json
     */
    /**
     * Sample code: List RAI bindings.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listRAIBindings(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBindings().list("resource-group", "safety-account", 50, com.azure.core.util.Context.NONE);
    }
}
