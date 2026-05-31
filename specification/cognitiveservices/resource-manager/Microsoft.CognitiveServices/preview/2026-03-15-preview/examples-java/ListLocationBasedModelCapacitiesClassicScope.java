
/**
 * Samples for LocationBasedModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListLocationBasedModelCapacitiesClassicScope.json
     */
    /**
     * Sample code: ListLocationBasedModelCapacities Classic Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listLocationBasedModelCapacitiesClassicScope(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.locationBasedModelCapacities().list("WestUS", "OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
