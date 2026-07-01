
/**
 * Samples for Computes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListComputes.json
     */
    /**
     * Sample code: ListComputes.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listComputes(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().list("rgcognitiveservices", "myAccount", com.azure.core.util.Context.NONE);
    }
}
