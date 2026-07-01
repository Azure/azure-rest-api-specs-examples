
/**
 * Samples for Computes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/DeleteCompute.json
     */
    /**
     * Sample code: DeleteCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().delete("rgcognitiveservices", "myAccount", "myCompute", com.azure.core.util.Context.NONE);
    }
}
