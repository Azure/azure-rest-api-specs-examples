
/**
 * Samples for Models List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * ListLocationModels.json
     */
    /**
     * Sample code: ListLocationModels.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listLocationModels(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.models().list("WestUS", com.azure.core.util.Context.NONE);
    }
}
