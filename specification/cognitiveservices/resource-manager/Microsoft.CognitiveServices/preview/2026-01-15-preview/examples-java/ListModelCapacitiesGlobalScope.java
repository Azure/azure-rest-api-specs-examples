
/**
 * Samples for ModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ListModelCapacitiesGlobalScope.json
     */
    /**
     * Sample code: ListModelCapacities Global Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listModelCapacitiesGlobalScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.modelCapacities().list("OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
