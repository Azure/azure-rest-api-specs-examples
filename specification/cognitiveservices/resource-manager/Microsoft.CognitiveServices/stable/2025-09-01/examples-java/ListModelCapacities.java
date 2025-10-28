
/**
 * Samples for ModelCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListModelCapacities.json
     */
    /**
     * Sample code: ListModelCapacities.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listModelCapacities(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.modelCapacities().list("OpenAI", "ada", "1", com.azure.core.util.Context.NONE);
    }
}
