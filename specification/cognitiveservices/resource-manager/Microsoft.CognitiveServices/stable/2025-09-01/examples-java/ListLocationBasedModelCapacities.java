
/**
 * Samples for LocationBasedModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListLocationBasedModelCapacities.json
     */
    /**
     * Sample code: ListLocationBasedModelCapacities.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listLocationBasedModelCapacities(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.locationBasedModelCapacities().list("WestUS", "OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
