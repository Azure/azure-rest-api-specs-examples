
/**
 * Samples for Workbenches List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListWorkbenches.json
     */
    /**
     * Sample code: ListWorkbenches.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listWorkbenches(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().list("rgcognitiveservices", "myAccount", "myProject", com.azure.core.util.Context.NONE);
    }
}
