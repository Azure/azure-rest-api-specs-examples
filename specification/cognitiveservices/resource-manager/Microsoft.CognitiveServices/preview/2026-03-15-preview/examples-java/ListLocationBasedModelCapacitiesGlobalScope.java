
/**
 * Samples for LocationBasedModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListLocationBasedModelCapacitiesGlobalScope.json
     */
    /**
     * Sample code: ListLocationBasedModelCapacities Global Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listLocationBasedModelCapacitiesGlobalScope(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.locationBasedModelCapacities().list("WestUS", "OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
