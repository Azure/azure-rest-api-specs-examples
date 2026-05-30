
/**
 * Samples for LocationBasedModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListLocationBasedModelCapacitiesDataZoneScope.json
     */
    /**
     * Sample code: ListLocationBasedModelCapacities DataZone Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listLocationBasedModelCapacitiesDataZoneScope(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.locationBasedModelCapacities().list("WestUS", "OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
