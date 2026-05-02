
/**
 * Samples for ModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ListModelCapacitiesDataZoneScope.json
     */
    /**
     * Sample code: ListModelCapacities DataZone Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listModelCapacitiesDataZoneScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.modelCapacities().list("OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
